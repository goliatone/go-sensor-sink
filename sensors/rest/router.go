package rest

import (
	"sensors"
	"sensors/registry"
	"sensors/rest/api/authentication"
	"sensors/rest/api/devices"
	"sensors/rest/api/middleware"
	"sensors/sink"

	"github.com/gofiber/fiber"
	fibermiddleware "github.com/gofiber/fiber/middleware"
)

//Router exposes the REST router to register our routes with the fiber app
func Router(app *fiber.App, domain *registry.Domain, config sensors.Config) {

	apiGroup := app.Group("/api")
	apiGroup.Use(fibermiddleware.Logger())

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
	// Readings
	////////////////////////////////////////////////////////////
	sinkRepo := domain.Readings
	// sinkRepo := sink.NewRepository(db)

	apiGroup.Post("/sensors/reading", func(c *fiber.Ctx) {
		var reading sink.DHT22Reading
		if err := c.BodyParser(&reading); err != nil {
			c.Status(503).Send(err)
			return
		}

		reading, err := sinkRepo.Add(reading)
		if err != nil {
			c.Status(503).Send(err)
			return
		}
		c.JSON(reading)
	})

	apiGroup.Get("/sensors/readings", func(c *fiber.Ctx) {
		qs := sink.NewSearchParameters(c)

		readings, err := sinkRepo.Get(qs)
		if err != nil {
			c.Status(503).Send(err)
			return
		}
		c.JSON(readings)
	})

	apiGroup.Get("/sensors/readings/:bucket", func(c *fiber.Ctx) {
		bucket := c.Params("bucket")

		items, err := sinkRepo.GetAggregateByBucket(bucket)
		if err != nil {
			c.Status(503).Send(err)
			return
		}
		c.JSON(items)
	})
}
