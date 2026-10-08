package patients

import "github.com/gofiber/fiber/v2"

// Module holds patients dependencies.
type Module struct{}

// New creates the patients module.
func New() *Module {
	return &Module{}
}

// Register mounts patients routes.
func (m *Module) Register(r fiber.Router) {
	r.Get("/patients/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "patients", "status": "ok"})
	})
}
