package rest

import (
	"sensors/device"
	"sensors/sink"
	"sensors/storage"

	"github.com/gofiber/fiber"
	fibermiddleware "github.com/gofiber/fiber/middleware"
	uuid "github.com/satori/go.uuid"
)

//Router exposes the REST router to register our routes with the fiber app
func Router(fiberApp *fiber.App, db *storage.Database) {

	apiGroup := fiberApp.Group("/api")
	apiGroup.Use(fibermiddleware.Logger())

	apiRouteGroup(apiGroup, db)
}

func apiRouteGroup(g *fiber.Group, db *storage.Database) {
	// g.Post("/login")
	deviceRepo := device.NewRepository(db)

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
	sinkRepo := sink.NewRepository(db)

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
}
