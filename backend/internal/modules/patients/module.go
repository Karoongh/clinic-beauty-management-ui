package patients

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
)

// Module holds patients dependencies.
type Module struct {
	repo *Repository
}

// New creates the patients module with a repository.
func New(repo *Repository) *Module {
	return &Module{repo: repo}
}

// Register mounts patients routes (protected by JWT).
func (m *Module) Register(r fiber.Router) {
	r.Get("/patients/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "patients", "status": "ok"})
	})

	protected := r.Group("/patients", auth.Middleware())
	protected.Get("/", m.list)
	protected.Get("/:id", m.get)
	protected.Get("/:id/wallet", m.wallet)
}
