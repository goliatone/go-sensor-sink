package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"sensors"
	"sensors/config"
	"sensors/event"
	"sensors/pubsub"
	"sensors/realtime"
	"sensors/rest"
	"sensors/sink"
	"sensors/storage/postgres"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber"
)

const (
	publicPath = "../frontend/src/build"
)

func main() {

	yaml := config.ReadYaml("")
	cnf := sensors.GetConfig(*yaml)

	log.Printf("mqtt %+v", cnf.Mqtt)

	database, err := postgres.NewDatabase(cnf)
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	server := fiber.New()
	emitter := event.NewEmitter()

	rest.Router(server, database)
	postgres.Migrate(database)
	realtime.Websockets(server)

	server.Static("/", publicPath)

	//move to package
	sinkRepo := sink.NewRepository(database)
	pubsub.SetConfig(&cnf.Mqtt)

	mqttClient := pubsub.AddCommandHandler(0, func(mc mqtt.Client, msg mqtt.Message) {
		//TODO: We get all messages, we should actually prefix it with server id so
		//that we don't get our own messages...
		// log.Printf("Reading message: %s %s", msg.Topic(), msg.Payload())
		if strings.Contains(msg.Topic(), "/reading") == false {
			return
		}

		reading, err := sink.NewDHT22Reading(msg.Payload())
		if err != nil {
			log.Println("error handling reading:" + err.Error())
			return
		}

		emitter.EmitSync("mqtt.event", reading)
	})

	if mqttClient != nil {
		log.Println("started mqtt")
	}

	emitter.On("mqtt.event", func(args ...interface{}) {
		reading := args[0].(sink.DHT22Reading)
		_, err = sinkRepo.Add(reading)

		if err != nil {
			log.Println("error adding reading:" + err.Error())
		}

		if message, err := reading.Deserialize(); err == nil {
			// wsm.Broadcast(message)
			realtime.Broadcast(message)
		}
	})

	port := fmt.Sprintf(":%s", cnf.Server.Port)
	log.Fatal(server.Listen(port))

	fmt.Sprintf("app key: %s\n", cnf.Secret)
}
