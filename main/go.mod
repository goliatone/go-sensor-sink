module github.com/goliatone/go-sensor-sink/main

go 1.13.3

require sensors v1.0.0

replace sensors => ../sensors

require (
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/gofiber/fiber v1.13.3
)
