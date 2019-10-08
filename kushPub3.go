package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
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

func main() {
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
	tlsconfig := NewTLSConfig()

	//testJson()
	opts := MQTT.NewClientOptions()
	opts.AddBroker("ssl://a2x45klxf2ndn6-ats.iot.ap-southeast-2.amazonaws.com:8883")
	opts.SetClientID(clientID + groupConfig + "Master").SetTLSConfig(tlsconfig)
	opts.SetDefaultPublishHandler(f)

	var ipPtr = flag.String("lightStatus", "foo", "a string")
	var count = flag.Int("numCounter", 1, "an int")
	flag.Parse()
	value := *ipPtr
	counterVar := *count

	// Start the connection
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		if groupConfig == "group1" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group1":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			// fmt.Println(sequence)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group2" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group2":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group3" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group3":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group4" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group4":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group5" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group5":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group6" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group6":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group7" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group7":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group8" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group8":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group9" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group9":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group10" && value == "on" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group10":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group1" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group1":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group2" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group2":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group3" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group3":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group4" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group4":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group5" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group5":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group6" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group6":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group7" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group7":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group8" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group8":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group9" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group9":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group10" && value == "off" && counterVar != 12 {
			shadowOn := `{"state":{"reported":{"group10":{"challenge3":"off"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Disconnect(250)
			fmt.Printf("\nMeesage has been published %v time\n", counterVar)
		} else if groupConfig == "group1" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group1":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group1":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group1":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
			os.Exit(1)
		} else if groupConfig == "group2" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group2":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group2":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group2":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group3" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group3":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group3":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group3":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group4" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group4":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group4":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group4":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group5" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group5":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group5":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group5":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group6" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group6":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group6":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group6":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group7" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group7":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group7":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group7":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group8" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group8":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group8":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group8":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group9" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group9":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group9":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group9":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else if groupConfig == "group10" && counterVar == 12 {
			shadowOn := `{"state":{"reported":{"group10":{"challenge3":"on"}}}}`
			c.Publish("$aws/things/demo/shadow/update", 1, false, shadowOn)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group10":{"challenge1":"on"}}}}`)
			c.Publish("$aws/things/demo/shadow/update", 1, false, `{"state":{"reported":{"group10":{"challenge2":"on"}}}}`)
			c.Disconnect(250)
		} else {
			panic("\nCouldn't allocate correct Group ID")

		}
	}
}
