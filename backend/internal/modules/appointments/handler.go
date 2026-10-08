package appointments

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	mu           sync.Mutex
	demoRequests = []Request{
		{ID: "r1", Channel: "telegram", Type: "reserve", Status: "pending", PatientName: "نگار احمدی", Mobile: "09123334455", Service: "بوتاکس", PreferredAt: "فردا صبح", CreatedAt: "2026-10-08T08:10:00Z"},
		{ID: "r2", Channel: "instagram", Type: "deposit", Status: "pending", PatientName: "لیلا موسوی", Mobile: "09351112233", Amount: 2000000, Notes: "بیعانه جلسه اول", CreatedAt: "2026-10-08T08:45:00Z"},
		{ID: "r3", Channel: "phone", Type: "reserve", Status: "accepted", PatientName: "حسین نوری", Service: "فیلر لب", PreferredAt: "امروز ۱۷", CreatedAt: "2026-10-07T16:00:00Z", DecidedAt: "2026-10-07T16:20:00Z", DecidedBy: "u-reception"},
		{ID: "r4", Channel: "whatsapp", Type: "reserve", Status: "rejected", PatientName: "مینا کاظمی", Service: "مشاوره", CreatedAt: "2026-10-07T11:00:00Z", DecidedAt: "2026-10-07T11:30:00Z", DecidedBy: "u-manager"},
		{ID: "r5", Channel: "telegram", Type: "reserve", Status: "pending", PatientName: "پارسا جعفری", Mobile: "09125556677", Service: "جوانسازی", PreferredAt: "هفته آینده", CreatedAt: "2026-10-08T09:00:00Z"},
	}
)

func (m *Module) list(c *fiber.Ctx) error {
	channel := c.Query("channel") // empty = all
	status := c.Query("status")   // empty = all
	typ := c.Query("type")       // empty = all

	result := make([]Request, 0, len(demoRequests))
	for _, r := range demoRequests {
		if channel != "" && r.Channel != channel {
			continue
		}
		if status != "" && r.Status != status {
			continue
		}
		if typ != "" && r.Type != typ {
			continue
		}
		result = append(result, r)
	}
	return c.JSON(fiber.Map{
		"data":  result,
		"total": len(result),
	})
}

func (m *Module) decide(c *fiber.Ctx) error {
	id := c.Params("id")
	var req DecideRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if req.Action != "accept" && req.Action != "reject" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "action must be accept or reject"})
	}

	mu.Lock()
	defer mu.Unlock()

	for i := range demoRequests {
		if demoRequests[i].ID == id {
			if demoRequests[i].Status != "pending" {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "request already decided"})
			}
			if req.Action == "accept" {
				demoRequests[i].Status = "accepted"
			} else {
				demoRequests[i].Status = "rejected"
			}
			demoRequests[i].DecidedAt = time.Now().UTC().Format(time.RFC3339)
			if uid, ok := c.Locals("userID").(string); ok {
				demoRequests[i].DecidedBy = uid
			}
			if req.Notes != "" {
				demoRequests[i].Notes = req.Notes
			}
			// Note: deposit accept → wallet credit will be handled later under accounting gate
			return c.JSON(demoRequests[i])
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "request not found"})
}

func (m *Module) get(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, r := range demoRequests {
		if r.ID == id {
			return c.JSON(r)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "request not found"})
}
