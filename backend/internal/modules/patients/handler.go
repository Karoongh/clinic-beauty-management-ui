package patients

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (m *Module) list(c *fiber.Ctx) error {
	if m.repo == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "patients repository not configured"})
	}
	list, err := m.repo.List(c.Context(), c.Query("q"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to list patients"})
	}
	if list == nil {
		list = []Patient{}
	}
	return c.JSON(fiber.Map{"data": list, "total": len(list)})
}

func (m *Module) get(c *fiber.Ctx) error {
	if m.repo == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "patients repository not configured"})
	}
	p, err := m.repo.GetByID(c.Context(), c.Params("id"))
	if err != nil {
		if errors.Is(err, errPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "patient not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to load patient"})
	}
	return c.JSON(p)
}

func (m *Module) wallet(c *fiber.Ctx) error {
	if m.repo == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "patients repository not configured"})
	}
	balance, err := m.repo.WalletBalance(c.Context(), c.Params("id"))
	if err != nil {
		if errors.Is(err, errPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "patient not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to load wallet"})
	}
	return c.JSON(fiber.Map{
		"patient_id": c.Params("id"),
		"balance":    balance,
		"currency":   "IRR",
	})
}
