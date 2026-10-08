package auth

import "github.com/gofiber/fiber/v2"

// Module holds auth dependencies.
type Module struct {
	repo *Repository
}

// New creates the auth module with a user repository.
func New(repo *Repository) *Module {
	return &Module{repo: repo}
}

// Register mounts auth routes under the given router.
func (m *Module) Register(r fiber.Router) {
	r.Post("/auth/login", m.login)

	protected := r.Group("/auth", Middleware())
	protected.Get("/me", m.me)

	r.Get("/auth/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"module": "auth", "status": "ok"})
	})
}
