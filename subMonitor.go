package main

import (
	"challenge1Monitor/kushpub"
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

// init RootCA, DeviceCert and DevicePrivateKey
func NewTLSConfig(rootCAPath, certificatePath, privateKeyPath string) *tls.Config {

	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(rootCAPath) //Amazon RootCA
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair(certificatePath, privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	// Just to print out the client certificate..
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

// init sub message callback
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Type is: %T\n", msg.Payload())
	fmt.Printf("Payload is: %s\n", msg.Payload())
	data := byteSlice(msg.Payload())
	if string(data) == "test" {
		fmt.Println("!!!")
	}
}

func main() {
	// init channel
	choke := make(chan [2]string)

	// init debug message
	MQTT.ERROR = log.New(os.Stdout, "", 0)

	// read file
	data, err := ioutil.ReadFile("/home/awstechevent/scripts/mqttclient_config.json")
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
	topic := obj.TOPIC

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
		//<-choke
		//fmt.Printf("%T\n", incoming[1])
		fmt.Println("\n--------Listening to your publishing now...--------")
		fmt.Printf("\nCaptured message from topic: %s\n", incoming[0])
		//messageGot := incoming[1]
		fmt.Println("\n----------------")
		fmt.Println("Processing to change the light status, please check the UI")

		kushpub.MaterPub()

		if err != nil {
			panic(err)
		}
		fmt.Println("\n----------------")
		fmt.Printf("Message Publish Successfully!!!\n")
		fmt.Println("Check the UI board, see if the light is on --:)")
		fmt.Print("\n\n-----Ending-----\n\n")
		break
	}
}
