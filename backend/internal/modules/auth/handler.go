package auth

import (
	"github.com/gofiber/fiber/v2"
)

// demoUser is temporary until the users table and repository are ready.
type demoUser struct {
	ID          string
	Username    string
	Password    string // plain only for local demo – never in production
	Role        string
	DisplayName string
}

// Local demo accounts. Replace with database + bcrypt as soon as users table exists.
var demoUsers = []demoUser{
	{ID: "u-manager", Username: "manager", Password: "manager123", Role: "manager", DisplayName: "مدیر"},
	{ID: "u-reception", Username: "reception", Password: "reception123", Role: "reception", DisplayName: "پذیرش"},
	{ID: "u-doctor", Username: "doctor", Password: "doctor123", Role: "doctor", DisplayName: "پزشک"},
	{ID: "u-cashier", Username: "cashier", Password: "cashier123", Role: "cashier", DisplayName: "صندوقدار"},
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
	if found == nil || found.Password != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
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
