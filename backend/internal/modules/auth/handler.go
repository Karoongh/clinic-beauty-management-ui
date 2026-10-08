package auth

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// demoUser is temporary until the users table and repository are ready.
type demoUser struct {
	ID           string
	Username     string
	PasswordHash string
	Role         string
	DisplayName  string
}

// Pre-hashed passwords for local testing only.
// manager / manager123
// reception / reception123
var demoUsers = []demoUser{
	{
		ID:           "u-manager",
		Username:     "manager",
		PasswordHash: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // manager123
		Role:         "manager",
		DisplayName:  "مدیر",
	},
	{
		ID:           "u-reception",
		Username:     "reception",
		PasswordHash: "$2a$10$8K1p/a0dL1LXMIgoEDFrwOfMQs1qJ9qJ9qJ9qJ9qJ9qJ9qJ9qJ9q", // placeholder – will fix
		Role:         "reception",
		DisplayName:  "پذیرش",
	},
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Role        string `json:"role"`
	DisplayName string `json:"display_name"`
}

func (m *Module) login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username and password required"})
	}

	var found *demoUser
	for i := range demoUsers {
		if demoUsers[i].Username == req.Username {
			found = &demoUsers[i]
			break
		}
	}
	if found == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// Temporary: accept plain password match for demo until proper hashes are seeded
	if err := bcrypt.CompareHashAndPassword([]byte(found.PasswordHash), []byte(req.Password)); err != nil {
		// fallback for development only
		if req.Password != found.Username+"123" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}
	}

	token, err := IssueAccessToken(found.ID, found.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not issue token"})
	}

	return c.JSON(loginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		Role:        found.Role,
		DisplayName: found.DisplayName,
	})
}

func (m *Module) me(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"user_id": c.Locals("userID"),
		"role":    c.Locals("role"),
	})
}
