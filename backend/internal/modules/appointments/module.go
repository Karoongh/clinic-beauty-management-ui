package appointments

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
)

// Module holds appointments dependencies.
type Module struct{}

// New creates the appointments module.
func New() *Module {
	return &Module{}
}

// Register mounts appointments routes.
func (m *Module) Register(r fiber.Router) {
	r.Get("/appointments/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "appointments", "status": "ok"})
	})

	protected := r.Group("/appointments", auth.Middleware())
	protected.Get("/requests", m.list)             // GET /api/v1/appointments/requests?channel=&status=&type=
	protected.Get("/requests/:id", m.get)          // GET /api/v1/appointments/requests/:id
	protected.Post("/requests/:id/decide", m.decide) // POST /api/v1/appointments/requests/:id/decide
}
