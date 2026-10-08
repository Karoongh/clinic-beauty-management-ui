package admissions

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
)

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

	protected := r.Group("/admissions", auth.Middleware())
	protected.Get("/", m.list)       // GET /api/v1/admissions
	protected.Get("/today", m.today) // GET /api/v1/admissions/today
	protected.Post("/", m.create)    // POST /api/v1/admissions
	protected.Get("/:id", m.get)     // GET /api/v1/admissions/:id
}
