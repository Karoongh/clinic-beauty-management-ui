package admissions

import "github.com/gofiber/fiber/v2"

// Module holds admissions dependencies.
type Module struct{}

// New creates the admissions module.
func New() *Module {
	return &Module{}
}

// Register mounts admissions routes.
func (m *Module) Register(r fiber.Router) {
	r.Get("/admissions/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "admissions", "status": "ok"})
	})
}
