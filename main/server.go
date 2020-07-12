package main

import (
	"fmt"
	"log"
	"os"

	"sensors"
	"sensors/config"
	"sensors/rest"
	"sensors/storage/postgres"

	"github.com/gofiber/fiber"
)

func main() {

	cfg := config.ReadYaml("")
	confg := sensors.GetConfig(*cfg)

	database, err := postgres.NewDatabase(confg)
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	postgres.Migrate(database)

	server := fiber.New()

	rest.Router(server, database)

	port := fmt.Sprintf(":%s", confg.Server.Port)
	log.Fatal(server.Listen(port))

	fmt.Sprintf("app key: %s\n", confg.Secret)
}
