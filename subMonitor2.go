package main

import (
	"challenge2Monitor/kushpub"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

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

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Type is: %T\n", msg.Payload())
	fmt.Printf("Payload is: %s\n", msg.Payload())
	data := byteSlice(msg.Payload())
	if string(data) == "test" {
		fmt.Println("!!!")
	}
}

func main() {
	choke := make(chan [2]string)

	MQTT.ERROR = log.New(os.Stdout, "", 0)
	// read file
	// data, err := ioutil.ReadFile("/home/awstechevent/scripts/mqttclient_config.json")
	data, err := ioutil.ReadFile("/Users/jswu/Desktop/Go/src/clientconfig.json")
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

	group := obj.GROUP
	endPoint := obj.ENDPOINT
	clientID := obj.CLIENTID
	rootCAPath := obj.ROOTCAPATH
	certificatePath := obj.CERTIFICATEPATH
	privateKeyPath := obj.PRIVATEKEYPATH
	port := obj.PORT
	tempTopic := obj.TOPIC
	topic := tempTopic + "/accepted"

	tlsconfig := NewTLSConfig(rootCAPath, certificatePath, privateKeyPath)

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
	for {
		incoming := <-choke

		fmt.Println("\n----------------")
		fmt.Println("Processing to change the light status, please check the UI")

		//make map for store response data
		shadowMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(incoming[1]), &shadowMap)
		if err != nil {
			//fmt.Println(`\nPlease make sure your shadow pub payload is: {"state":{"reported":{"groupX":{"challenge2":"on/off"}}}}`)
			panic(err)
		}

		//parse response mapped data
		StateMap := shadowMap["state"].(map[string]interface{})

		ReportedMap := StateMap["reported"].(map[string]interface{})

		GroupMap := ReportedMap[group].(map[string]interface{})
		Light := GroupMap["challenge2"]
		fmt.Println(Light)

		if Light == "on" {
			kushpub.MaterPub()
			fmt.Println("\n----------------")
			fmt.Printf("Shadow has now been succefully updated\n")
			fmt.Print("\n\n-----Ending-----\n\n")
			break
		} else if Light == "off" {
			panic("Seems like you get incorrect Light status")

		}
	}
}
