package patients

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
)

// Module holds patients dependencies.
type Module struct{}

// New creates the patients module.
func New() *Module {
	return &Module{}
}

// Register mounts patients routes (all protected by JWT).
func (m *Module) Register(r fiber.Router) {
	// Smoke test (public)
	r.Get("/patients/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "patients", "status": "ok"})
	})

	// Protected routes
	protected := r.Group("/patients", auth.Middleware())
	protected.Get("/", m.list)           // GET /api/v1/patients?q=
	protected.Get("/:id", m.get)         // GET /api/v1/patients/:id
	protected.Get("/:id/wallet", m.wallet) // GET /api/v1/patients/:id/wallet
}
