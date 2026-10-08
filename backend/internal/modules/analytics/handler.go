package analytics

import "github.com/gofiber/fiber/v2"

// homeKPIs returns the main dashboard numbers for the Home screen.
// Values are demo aggregates until real repositories exist.
func (m *Module) homeKPIs(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"admissions_today":    3,
		"pending_requests":    3,
		"patients_total":      4,
		"revenue_today":       12500000, // Rials – display only
		"wallet_credit_total": 2000000,  // Rials – display only
		"currency":            "IRR",
		"note":                "demo aggregates – replace with real queries later",
	})
}
