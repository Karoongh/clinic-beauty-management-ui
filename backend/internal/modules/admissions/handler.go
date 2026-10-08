package admissions

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	mu            sync.Mutex
	nextNumber    = 5 // next admission number (demo starts after existing ones)
	demoAdmissions = []Admission{
		{ID: "a1", Number: 1, Date: "2026-10-08", PatientID: "p1", PatientName: "سارا محمدی", Channel: "telegram", Service: "بوتاکس", Doctor: "دکتر احمدی", CreatedAt: "2026-10-08T09:15:00Z"},
		{ID: "a2", Number: 2, Date: "2026-10-08", PatientID: "p3", PatientName: "مریم حسینی", Channel: "walk-in", Service: "فیلر", Doctor: "دکتر احمدی", CreatedAt: "2026-10-08T10:00:00Z"},
		{ID: "a3", Number: 3, Date: "2026-10-07", PatientID: "p2", PatientName: "علی رضایی", Channel: "phone", Service: "مشاوره", CreatedAt: "2026-10-07T14:30:00Z"},
		{ID: "a4", Number: 4, Date: "2026-10-08", PatientID: "p4", PatientName: "رضا کریمی", Channel: "instagram", Service: "جوانسازی", Doctor: "دکتر حسینی", CreatedAt: "2026-10-08T11:20:00Z"},
	}
	validChannels = map[string]bool{
		"walk-in": true, "phone": true, "telegram": true, "instagram": true, "whatsapp": true,
	}
)

func (m *Module) list(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data":  demoAdmissions,
		"total": len(demoAdmissions),
	})
}

func (m *Module) today(c *fiber.Ctx) error {
	today := time.Now().Format("2006-01-02")
	result := make([]Admission, 0)
	for _, a := range demoAdmissions {
		if a.Date == today {
			result = append(result, a)
		}
	}
	return c.JSON(fiber.Map{
		"data":  result,
		"total": len(result),
		"date":  today,
	})
}

func (m *Module) create(c *fiber.Ctx) error {
	var req CreateAdmissionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if req.PatientID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "patient_id is required"})
	}
	if !validChannels[req.Channel] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "channel must be one of: walk-in, phone, telegram, instagram, whatsapp",
		})
	}

	date := req.Date
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	mu.Lock()
	num := nextNumber
	nextNumber++
	mu.Unlock()

	a := Admission{
		ID:          fmt.Sprintf("a%d", num),
		Number:      num,
		Date:        date,
		PatientID:   req.PatientID,
		PatientName: "", // will be filled from patients module later
		Channel:     req.Channel,
		Service:     req.Service,
		Doctor:      req.Doctor,
		Notes:       req.Notes,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
	}

	mu.Lock()
	demoAdmissions = append(demoAdmissions, a)
	mu.Unlock()

	return c.Status(fiber.StatusCreated).JSON(a)
}

func (m *Module) get(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, a := range demoAdmissions {
		if a.ID == id {
			return c.JSON(a)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "admission not found"})
}
