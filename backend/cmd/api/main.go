package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/config"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/admissions"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/analytics"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/appointments"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/catalog"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/patients"
	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/platform"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := platform.NewDBPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer pool.Close()

	app := fiber.New(fiber.Config{
		AppName: "Clinic Beauty API",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	v1 := app.Group("/api/v1")

	authRepo := auth.NewRepository(pool)
	patientsRepo := patients.NewRepository(pool)

	auth.New(authRepo).Register(v1)
	patients.New(patientsRepo).Register(v1)
	admissions.New().Register(v1)
	appointments.New().Register(v1)
	catalog.New().Register(v1)
	analytics.New().Register(v1)

	// Finance / inventory remain gated by accounting-audit checklist

	addr := ":" + cfg.Port
	log.Printf("listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
