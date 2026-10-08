package auth

import "github.com/gofiber/fiber/v2"

// Module holds auth dependencies (to be injected later).
type Module struct{}

// New creates the auth module.
func New() *Module {
	return &Module{}
}

// Register mounts auth routes under the given router.
// Currently only a placeholder.
func (m *Module) Register(r fiber.Router) {
	r.Get("/auth/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "auth", "status": "ok"})
	})
}
