package patients

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// demo data – replace with repository + PostgreSQL later
var demoPatients = []Patient{
	{ID: "p1", Code: "P-1001", FullName: "سارا محمدی", Mobile: "09121234567", Age: 32, Gender: "female", Visits: 5, Invoices: 4, TotalPaid: 12500000, Wallet: 500000, Status: "active"},
	{ID: "p2", Code: "P-1002", FullName: "علی رضایی", Mobile: "09129876543", Age: 41, Gender: "male", Visits: 2, Invoices: 2, TotalPaid: 4800000, Wallet: 0, Status: "active"},
	{ID: "p3", Code: "P-1003", FullName: "مریم حسینی", Mobile: "09351234567", Age: 28, Gender: "female", Visits: 8, Invoices: 7, TotalPaid: 21000000, Wallet: 1500000, Status: "active"},
	{ID: "p4", Code: "P-1004", FullName: "رضا کریمی", Mobile: "09121112233", Age: 35, Gender: "male", Visits: 1, Invoices: 1, TotalPaid: 2500000, Wallet: 0, Status: "active"},
}

func (m *Module) list(c *fiber.Ctx) error {
	q := strings.TrimSpace(strings.ToLower(c.Query("q")))
	result := make([]Patient, 0, len(demoPatients))
	for _, p := range demoPatients {
		if q == "" ||
			strings.Contains(strings.ToLower(p.FullName), q) ||
			strings.Contains(p.Mobile, q) ||
			strings.Contains(strings.ToLower(p.Code), q) {
			result = append(result, p)
		}
	}
	return c.JSON(fiber.Map{
		"data":  result,
		"total": len(result),
	})
}

func (m *Module) get(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, p := range demoPatients {
		if p.ID == id {
			return c.JSON(PatientProfile{
				Patient: p,
				Notes:   "",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "patient not found"})
}

func (m *Module) wallet(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, p := range demoPatients {
		if p.ID == id {
			return c.JSON(fiber.Map{
				"patient_id": p.ID,
				"balance":    p.Wallet,
				"currency":   "IRR",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "patient not found"})
}
