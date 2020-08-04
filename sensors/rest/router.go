package rest

import (
	"sensors"
	"sensors/registry"
	"sensors/rest/api/authentication"
	"sensors/rest/api/devices"
	"sensors/rest/api/middleware"
	"sensors/rest/api/readings"
	"sensors/rest/api/users"

	"github.com/gofiber/fiber"
	// fibermiddleware "github.com/gofiber/fiber/middleware"
)

//Router exposes the REST router to register our routes with the fiber app
func Router(app *fiber.App, domain *registry.Domain, config sensors.Config) {

	apiGroup := app.Group("/api")
	//read from config sensors.Config.App.UseLogger
	// apiGroup.Use(fibermiddleware.Logger())

	apiGroup.Get("/status", func(ctx *fiber.Ctx) {
		ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
			"status": "online",
		})
	})

	authGroup := apiGroup.Group("/auth")
	authGroup.Post("/login", authentication.Login(domain.Auth))
	authGroup.Post("/user", authentication.Register(domain.Auth))
	// authGroup.Get("/user", )

	secret := config.Auth.JWTSecret
	authGroup.Post("/test", middleware.AuthByBearerToken(secret), func(ctx *fiber.Ctx) {
		response := map[string]interface{}{
			"success": true,
			"user":    ctx.Locals("user"),
		}

		ctx.JSON(response)
	})

	v1 := apiGroup.Group("/v1")
	v1.Use(middleware.AuthByBearerToken(secret))

	////////////////////////////////////////////////////////////
	// Device
	////////////////////////////////////////////////////////////

	v1.Post("/device", devices.Create(domain.Devices))
	v1.Get("/device/:id", devices.Read(domain.Devices))
	v1.Put("/device/:id", devices.Update(domain.Devices))
	v1.Delete("/device/:id", devices.Delete(domain.Devices))
	v1.Get("/device", devices.List(domain.Devices))

	////////////////////////////////////////////////////////////
	// User
	////////////////////////////////////////////////////////////

	v1.Post("/user", users.Create(domain.Users))
	v1.Get("/user/:id", users.Read(domain.Users))
	v1.Put("/user/:id", users.Update(domain.Users))
	v1.Delete("/user/:id", users.Delete(domain.Users))
	v1.Get("/user", users.List(domain.Users))

	////////////////////////////////////////////////////////////
	// Readings
	////////////////////////////////////////////////////////////

	v1.Post("/sensors/reading", readings.Create(domain.Readings))
	v1.Get("/sensors/readings", readings.List(domain.Readings))
	v1.Get("/sensors/readings/:bucket", readings.ListByBucket(domain.Readings))
}
