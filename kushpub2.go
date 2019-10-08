package kushpub

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func NewTLSConfig() *tls.Config {
	//init debug message
	MQTT.ERROR = log.New(os.Stdout, "", 0)

	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("/home/awstechevent/run/kushcerts/amazon.pem")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair("/home/awstechevent/run/kushcerts/device.crt", "/home/awstechevent/run/kushcerts/private.key")
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
}

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func MaterPub() {

	// init debug message
	MQTT.ERROR = log.New(os.Stdout, "", 0)
	// read file
	data, err := ioutil.ReadFile("/home/awstechevent/scripts/mqttclient_config.json")
	if err != nil {
		fmt.Print(err)
	}

	//group struct init
	type GroupConfig struct {
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
	var obj2 GroupConfig

	err = json.Unmarshal(data, &obj2)
	if err != nil {
		fmt.Println("error:", err)
	}
	clientID := obj2.CLIENTID
	groupConfig := obj2.GROUP
	fmt.Printf("Publishing to Group: %s\n", groupConfig)
	tlsconfig := NewTLSConfig()

	//testJson()
	opts := MQTT.NewClientOptions()
	opts.AddBroker("ssl://a2x45klxf2ndn6-ats.iot.ap-southeast-2.amazonaws.com:8883")
	opts.SetClientID(clientID + groupConfig + "Master").SetTLSConfig(tlsconfig)
	opts.SetDefaultPublishHandler(f)

	// Start the connection
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		if groupConfig == "group1" {
			shadowOn := `{"state":{"reported":{"group1":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group2" {
			shadowOn := `{"state":{"reported":{"group2":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group3" {
			shadowOn := `{"state":{"reported":{"group3":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group4" {
			shadowOn := `{"state":{"reported":{"group4":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group5" {
			shadowOn := `{"state":{"reported":{"group5":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group6" {
			shadowOn := `{"state":{"reported":{"group6":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group7" {
			shadowOn := `{"state":{"reported":{"group7":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group8" {
			shadowOn := `{"state":{"reported":{"group8":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group9" {
			shadowOn := `{"state":{"reported":{"group9":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else if groupConfig == "group10" {
			shadowOn := `{"state":{"reported":{"group10":{"challenge2":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Println("\nMeesage has been published")
		} else {
			panic("\nCouldn't allocate correct Group ID")
		}
	}
}
