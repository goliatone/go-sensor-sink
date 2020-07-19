package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	ikisocket "github.com/antoniodipinto/ikisocket"

	"sensors"
	"sensors/config"
	"sensors/pubsub"
	"sensors/rest"
	"sensors/sink"
	"sensors/storage/postgres"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber"
)

type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

func newReading(msg []byte) (sink.DHT22Reading, error) {
	reading := sink.DHT22Reading{}
	err := json.Unmarshal(msg, &reading)
	if err != nil {
		return sink.DHT22Reading{}, err
	}
	return reading, nil
}

func main() {

	yaml := config.ReadYaml("")
	cnf := sensors.GetConfig(*yaml)

	log.Printf("mqtt %+v", cnf.Mqtt)

	database, err := postgres.NewDatabase(cnf)
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	postgres.Migrate(database)

	server := fiber.New()

	server.Static("/", "../frontend/src/build")

	clients := make(map[string]string)

	server.Use(func(c *fiber.Ctx) {
		c.Locals("user_id", c.Query("user_id"))
		c.Next()
	})

	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		fmt.Println("ws connect")
	})

	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {
		fmt.Printf("ws message-> name %s data %s uuid %s \n", ep.Name, string(ep.Data), ep.SocketUUID)
	})

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		fmt.Println("fired disconnect" + ep.Error.Error())
	})

	server.Get("/ws", ikisocket.New(func(kws *ikisocket.Websocket) {
		userID := fmt.Sprintf("%v", kws.Locals("user_id"))
		fmt.Printf("ws user id: %s uuid: %s", userID, kws.UUID)

		kws.SetAttribute("user_id", userID)

		kws.OnConnect = func() {
			clients[userID] = kws.UUID
			kws.Emit([]byte("Hello user " + userID))
			kws.Broadcast([]byte("User connected "+userID+" UUID: "+kws.UUID), true)
		}

		kws.OnMessage = func(data []byte) {
			message := MessageObject{}
			json.Unmarshal(data, &message)

			fmt.Printf("message to %s: %s\n", message.To, message.Data)

			err := kws.EmitTo(clients[message.To], data)
			if err != nil {
				fmt.Printf("message %s error: %v\n", message.From, err)
			}
		}
	}))

	ikisocket.On("close", func(payload *ikisocket.EventPayload) {
		fmt.Println("ws closed " + payload.SocketAttributes["user_id"])
	})

	//move to package
	sinkRepo := sink.NewRepository(database)
	pubsub.SetConfig(&cnf.Mqtt)

	mqttClient := pubsub.AddCommandHandler(0, func(mc mqtt.Client, msg mqtt.Message) {
		//TODO: We get all messages, we should actually prefix it with server id so
		//that we don't get our own messages...
		// log.Printf("Reading message: %s %s", msg.Topic(), msg.Payload())
		if strings.Contains(msg.Topic(), "/readings") == false {
			return
		}

		reading, err := newReading(msg.Payload())
		if err != nil {
			log.Println("error handling reading:" + err.Error())
			return
		}
		_, err = sinkRepo.Add(reading)
		if err != nil {
			log.Println("error adding reading:" + err.Error())
		}
	})
	if mqttClient != nil {
		log.Println("started mqtt")
	}

	rest.Router(server, database)

	port := fmt.Sprintf(":%s", cnf.Server.Port)
	log.Fatal(server.Listen(port))

	fmt.Sprintf("app key: %s\n", cnf.Secret)
}
