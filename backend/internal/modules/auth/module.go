package auth

import "github.com/gofiber/fiber/v2"

// Module holds auth dependencies (DB repository will be injected later).
type Module struct{}

// New creates the auth module.
func New() *Module {
	return &Module{}
}

// Register mounts auth routes under the given router.
func (m *Module) Register(r fiber.Router) {
	// Public
	r.Post("/auth/login", m.login)

	// Protected
	protected := r.Group("/auth", Middleware())
	protected.Get("/me", m.me)

	// Keep ping for smoke tests
	r.Get("/auth/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "auth", "status": "ok"})
	})
}
