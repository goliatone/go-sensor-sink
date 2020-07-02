package rest

import (
	"github.com/gofiber/fiber"
	fibermiddleware "github.com/gofiber/fiber/middleware"
)

//Router exposes the REST router to register our routes with the fiber app
func Router(fiberApp *fiber.App) {
	apiGroup := fiberApp.Group("/api")
	apiGroup.Use(fibermiddleware.Logger())

	apiRouteGroup(apiGroup)
}

func apiRouteGroup(g *fiber.Group) {
	// g.Post("/login")

	g.Get("/sensors/readings", func(c *fiber.Ctx) {
		c.Send("Hi")
	})
}
