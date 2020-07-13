package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"sensors"
	"sensors/config"
	"sensors/pubsub"
	"sensors/rest"
	"sensors/sink"
	"sensors/storage/postgres"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber"
)

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

	//move to package
	sinkRepo := sink.NewRepository(database)
	pubsub.SetConfig(&cnf.Mqtt)
	mqttClient := pubsub.AddCommandHandler(0, func(mc mqtt.Client, msg mqtt.Message) {
		log.Printf("Reading message: %s %s", msg.Topic(), msg.Payload())
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
