module github.com/goliatone/go-sensor-sink/main

go 1.13.3

require sensors v1.0.0

replace sensors => ../sensors

require (
	github.com/antoniodipinto/ikisocket v0.0.0-20200526172531-f031f26ec81d // indirect
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/gofiber/fiber v1.12.2
)
