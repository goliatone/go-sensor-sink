package main

import (
	"fmt"
	"log"

	"sensors/config"
	"sensors/rest"

	"github.com/gofiber/fiber"
)

func main() {

	cfg := config.ReadYaml("")
	confg := sensors.GetConfig(*cfg)

	server := fiber.New()

	rest.Router(server)

	server.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World!")
	})

	port := fmt.Sprintf(":%v", 8372)
	log.Fatal(server.Listen(port))

	fmt.Sprintf("app key: %s\n", confg.Secret)
}
