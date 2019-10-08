package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	//Golang paho library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func NewTLSConfig(rootCAPath, certificatePath, privateKeyPath string) *tls.Config {

	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(rootCAPath)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	cert, err := tls.LoadX509KeyPair(certificatePath, privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{

		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
}

func byteSlice(b []byte) []byte { return b }

// Number increment
// var counter int
func increment(i *int) {
	*i = *i + 1
}

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Type is: %T\n", msg.Payload())
	fmt.Printf("Payload is: %s\n", msg.Payload())
	data := byteSlice(msg.Payload())
	if string(data) == "test" {
		fmt.Println("!!!")
	}
}

func main() {
	counter := 0
	choke := make(chan [2]string)

	MQTT.ERROR = log.New(os.Stdout, "", 0)
	// read file
	data, err := ioutil.ReadFile("/home/awstechevent/scripts/mqttclient_config.json")
	// data, err := ioutil.ReadFile("/Users/jswu/Desktop/Go/src/clientconfig.json")
	if err != nil {
		fmt.Print(err)
	}

	// define data structure
	type ClientConfig struct {
		GROUP           string
		ENDPOINT        string
		CLIENTID        string
		ROOTCAPATH      string
		CERTIFICATEPATH string
		PRIVATEKEYPATH  string
		PORT            int16
		TOPIC           string
	}

	// json data
	var obj ClientConfig

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}
	// can access using struct now
	fmt.Println("\n----------------")
	fmt.Println("Processing JSON Configuration ... ")
	fmt.Printf("\nGroup is : %s\n", obj.GROUP)
	fmt.Printf("EndPoint is : %s\n", obj.ENDPOINT)
	fmt.Printf("ClientId is : %s\n", obj.CLIENTID)
	fmt.Printf("RootCAPath is : %s\n", obj.ROOTCAPATH)
	fmt.Printf("DeviceCertPath is : %s\n", obj.CERTIFICATEPATH)
	fmt.Printf("DevicePrivatePath is : %s\n", obj.PRIVATEKEYPATH)
	fmt.Printf("Port is : %d\n", obj.PORT)
	fmt.Printf("Topic is : %s\n", obj.TOPIC)

	group := strings.ToLower(obj.GROUP)
	endPoint := strings.ToLower(obj.ENDPOINT)
	clientID := strings.ToLower(obj.CLIENTID)
	rootCAPath := strings.ToLower(obj.ROOTCAPATH)
	certificatePath := strings.ToLower(obj.CERTIFICATEPATH)
	privateKeyPath := strings.ToLower(obj.PRIVATEKEYPATH)
	port := obj.PORT
	tempTopic := strings.ToLower(obj.TOPIC)
	topic := strings.ToLower(tempTopic + "/accepted")

	tlsconfig := NewTLSConfig(rootCAPath, certificatePath, privateKeyPath)

	//MQTT client config
	opts := MQTT.NewClientOptions()
	endPointBuff := []byte("ssl://")
	endPointBuff = append(endPointBuff, endPoint...)
	endPointBuff = append(endPointBuff, ":"...)
	endPointBuff = append(endPointBuff, strconv.FormatInt((int64(port)), 10)...)
	sslendPoint := string(endPointBuff)
	fmt.Println(sslendPoint)
	opts.AddBroker(sslendPoint)
	opts.SetClientID(group + clientID).SetTLSConfig(tlsconfig)
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}

	})

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//checking correct sequence
	sequence := []string{}
	correctSequence := []string{"on", "off", "on", "off", "on", "off", "on", "off", "on", "off", "on", "off"}
	for {

		//listening to candidates pub topic
		incoming := <-choke
		increment(&counter)

		fmt.Println("\n*******************************")
		fmt.Println("Processing to change the light status, please check the UI")

		//make map for store response data
		shadowMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(incoming[1]), &shadowMap)
		if err != nil {
			panic(err)
		}

		//parse response mapped data
		StateMap := shadowMap["state"].(map[string]interface{})
		ReportedMap := StateMap["reported"].(map[string]interface{})
		GroupMap := ReportedMap[group].(map[string]interface{})
		Light := GroupMap["challenge3"]
		fmt.Printf("\nTrying to publish %v status . . .\n", Light)
		sequence = append(sequence, fmt.Sprintf("%s", Light))
		fmt.Println("\nYour current sequence is: ")
		fmt.Println(sequence)

		if Light == "on" && counter < 12 {

			var cmd = exec.Command("kushPub", "-lightStatus=on", fmt.Sprintf("-numCounter=%v", counter))
			output, err := cmd.Output()
			if err != nil {

				fmt.Println(err)
			}
			fmt.Printf("%v\n", string(output))

		} else if Light == "off" && counter < 12 {

			var cmd = exec.Command("kushPub", "-lightStatus=off", fmt.Sprintf("-numCounter=%v", counter))
			output, err := cmd.Output()
			if err != nil {

				fmt.Println(err)
			}
			fmt.Printf("%v\n", string(output))

		} else if Light == "on" && counter == 12 && reflect.DeepEqual(sequence, correctSequence) {
			var cmd = exec.Command("kushPub", "-lightStatus=on", fmt.Sprintf("-numCounter=%v", counter))
			output, err := cmd.Output()
			if err != nil {

				fmt.Println(err)
			}
			fmt.Printf("%v\n", string(output))
			fmt.Println("\nCongratulations!!! Misson Completed!!!")
			fmt.Println("\nDone! Light has been changed 12 times!!!")
			fmt.Println("\nPlease cehck the UI, all lights in your group should be all ON now")
			// os.Exit(1)
			break
		} else if Light == "off" && counter == 12 && reflect.DeepEqual(sequence, correctSequence) {
			var cmd = exec.Command("kushPub", "-lightStatus=on", fmt.Sprintf("-numCounter=%v", counter))
			output, err := cmd.Output()
			if err != nil {

				fmt.Println(err)
			}
			fmt.Printf("%v\n", string(output))
			fmt.Println("\nCongratulations!!! Misson Completed!!!")
			fmt.Println("\nDone! Light has been changed 12 times!!!")
			fmt.Println("\nPlease cehck the UI, all lights in your group should be ON now")
			// os.Exit(1)
			break
		} else {

			fmt.Println("\nX X X X X X X X X X X")
			fmt.Println("\nERROR: ")
			fmt.Println("\nPlease check your pub message payload, and meke sure they are in: [on, off, on, off, on, off, on, off, on, off, on, off] order")
			var cmd = exec.Command("kushPub", "-lightStatus=off", "-numCounter=13")
			output, err := cmd.Output()
			if err != nil {

				fmt.Println(err)
			}
			fmt.Printf("%v\n", string(output))
			//panic("Seems like you get incorrect Lighit status")
			os.Exit(1)
		}
	}
}
