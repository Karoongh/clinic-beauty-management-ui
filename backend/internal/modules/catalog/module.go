package catalog

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Karoongh/clinic-beauty-management-ui/backend/internal/modules/auth"
)

// Module holds catalog dependencies.
type Module struct{}

// New creates the catalog module.
func New() *Module {
	return &Module{}
}

// Register mounts catalog routes.
func (m *Module) Register(r fiber.Router) {
	r.Get("/catalog/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "catalog", "status": "ok"})
	})

	protected := r.Group("/catalog", auth.Middleware())
	protected.Get("/items", m.list)     // GET /api/v1/catalog/items?kind=&q=
	protected.Get("/items/:id", m.get)  // GET /api/v1/catalog/items/:id
}
