package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/admissions"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/appointments"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/catalog"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/patients"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Clinic Beauty API",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	// Health — used by Docker and load balancers
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// API v1 group
	v1 := app.Group("/api/v1")

	// Register feature modules (each module is isolated)
	auth.New().Register(v1)
	patients.New().Register(v1)
	admissions.New().Register(v1)
	appointments.New().Register(v1)
	catalog.New().Register(v1)

	// TODO: inventory, finance, analytics, settings
	// Finance module must wait for accounting-audit checklist completion

	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = ":" + v
	}

	log.Printf("listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
