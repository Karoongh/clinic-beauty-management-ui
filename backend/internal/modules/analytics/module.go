package analytics

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
)

// Module provides read-only dashboard metrics.
type Module struct{}

// New creates the analytics module.
func New() *Module {
	return &Module{}
}

// Register mounts analytics routes.
func (m *Module) Register(r fiber.Router) {
	r.Get("/analytics/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "analytics", "status": "ok"})
	})

	protected := r.Group("/analytics", auth.Middleware())
	protected.Get("/home", m.homeKPIs)
}
