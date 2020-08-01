package rest

import (
	"sensors/device"
	"sensors/registry"
	"sensors/rest/api/authentication"
	"sensors/sink"

	"github.com/gofiber/fiber"
	fibermiddleware "github.com/gofiber/fiber/middleware"
	uuid "github.com/satori/go.uuid"
)

//Router exposes the REST router to register our routes with the fiber app
func Router(app *fiber.App, domain *registry.Domain) {

	apiGroup := app.Group("/api")
	apiGroup.Use(fibermiddleware.Logger())

	authGroup := apiGroup.Group("/auth")
	authGroup.Post("/login", authentication.Login(domain.Auth))

	authGroup.Post("/user", authentication.Register(domain.Auth))

	apiRouteGroup(apiGroup.(*fiber.Group), domain)
}

func apiRouteGroup(g *fiber.Group, domain *registry.Domain) {
	// g.Post("/login")
	deviceRepo := domain.Devices
	////////////////////////////////////////////////////////////
	// Device
	////////////////////////////////////////////////////////////
	g.Get("/device", func(c *fiber.Ctx) {
		devices, err := deviceRepo.Get()
		if err != nil {
			c.Status(503).Send(err)
			return
		}
		c.JSON(devices)
	})

	g.Get("/device/:id", func(c *fiber.Ctx) {
		id, _ := uuid.FromString(c.Params("id"))

		item, err := deviceRepo.GetByID(id)
		if err != nil {
			res := make(map[string]interface{})
			res["error"] = true
			res["message"] = err.Error()
			c.JSON(res)
			return
		}
		c.JSON(item)
	})

	g.Post("/device", func(c *fiber.Ctx) {
		var item device.Device
		if err := c.BodyParser(&item); err != nil {
			c.Status(503).Send(err)
			return
		}

		d, err := deviceRepo.Add(item)
		if err != nil {
			c.Status(503).Send(err)
			return
		}

		c.JSON(d)
	})

	g.Put("/device/:id", func(c *fiber.Ctx) {
		id, _ := uuid.FromString(c.Params("id"))

		var item device.Device

		item, err := deviceRepo.GetByID(id)
		if err != nil {
			c.Status(503).Send(err)
			return
		}

		if err := c.BodyParser(&item); err != nil {
			c.Status(503).Send(err)
			return
		}

		if err := deviceRepo.Update(item); err != nil {
			c.Status(503).Send(err)
			return
		}

		c.JSON(item)
	})

	g.Delete("/device/:id", func(c *fiber.Ctx) {
		id, _ := uuid.FromString(c.Params("id"))

		err := deviceRepo.DeleteByID(id)
		if err != nil {
			res := make(map[string]interface{})
			res["error"] = true
			res["message"] = err.Error()
			c.JSON(res)
			return
		}
		c.Send("OK")
	})

	////////////////////////////////////////////////////////////
	// Readings
	////////////////////////////////////////////////////////////
	sinkRepo := domain.Readings
	// sinkRepo := sink.NewRepository(db)

	g.Post("/sensors/reading", func(c *fiber.Ctx) {
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

	g.Get("/sensors/readings", func(c *fiber.Ctx) {
		qs := sink.NewSearchParameters(c)

		readings, err := sinkRepo.Get(qs)
		if err != nil {
			c.Status(503).Send(err)
			return
		}
		c.JSON(readings)
	})

	g.Get("/sensors/readings/:bucket", func(c *fiber.Ctx) {
		bucket := c.Params("bucket")

		items, err := sinkRepo.GetAggregateByBucket(bucket)
		if err != nil {
			c.Status(503).Send(err)
			return
		}
		c.JSON(items)
	})
}
