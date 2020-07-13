package pubsub

import (
	"fmt"
	"log"
	"os"
	"sensors"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client
var clientID string
var cnf *sensors.Mqtt

type mqttClient struct {
	ID     string
	mqtt   mqtt.Client
	config *sensors.Mqtt
}

//SetConfig sets configuration for Mqtt client
func SetConfig(c *sensors.Mqtt) {
	cnf = c
	clientID = cnf.ClientID
	cnf.TopicInput = strings.Replace(cnf.TopicInput, ":client_id", clientID, -1)
	cnf.TopicOutput = strings.Replace(cnf.TopicOutput, ":client_id", clientID, -1)
}

//AddCommandHandler creates client add handlers
func AddCommandHandler(qos byte, callback mqtt.MessageHandler) mqtt.Client {
	hostname, _ := os.Hostname()

	if cnf.ClientID == "" {
		cnf.ClientID = strconv.Itoa(time.Now().Second()) + "-" + hostname
	}

	brokerEndpoint := fmt.Sprintf("%s://%s:%s", cnf.Scheme, cnf.Host, cnf.Port)
	log.Printf("mqtt endpoint %s\n", brokerEndpoint)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerEndpoint)
	opts.SetClientID(cnf.ClientID)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(false)

	//TODO: get from config
	retained := false
	topic := fmt.Sprintf("iot/device/%s/action/down", clientID)
	payload := fmt.Sprintf("{\"action\":\"down\", \"client\":\"%s\"}", clientID)
	opts.SetWill(topic, payload, qos, retained)

	if cnf.User != "" {
		opts.SetUsername(cnf.User)
	}

	if cnf.Password != "" {
		opts.SetPassword(cnf.Password)
	}

	// mqtt.WARN = logger.Warn
	// mqtt.CRITICAL = logger.Error
	// mqtt.ERROR = logger.Error

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panicln(token.Error())
	}

	log.Printf("subscribe: %s", cnf.TopicInput)
	client.Subscribe(cnf.TopicInput, qos, callback)

	topic = fmt.Sprintf("iot/device/%s/action/up", clientID)
	payload = fmt.Sprintf("{\"action\":\"up\", \"client\":\"%s\"}", clientID)

	if token := client.Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		log.Panicln(token.Error())
	}

	ping := time.NewTicker(25 * time.Second)
	go func(c mqtt.Client) {
		for {
			select {
			case <-ping.C:
				topic = fmt.Sprintf("iot/device/%s/action/ping", clientID)
				payload = fmt.Sprintf("{\"action\":\"ping\", \"client\":\"%s\"}", clientID)
				if token := c.Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
					log.Panicln(token.Error())
				}
			}
		}
	}(client)

	return client

}
