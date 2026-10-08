package catalog

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

var demoItems = []Item{
	{ID: "i1", Code: "S-001", Name: "بوتاکس پیشانی", Kind: "service", Unit: "واحد", Price: 3500000, Active: true},
	{ID: "i2", Code: "S-002", Name: "فیلر لب", Kind: "service", Unit: "سی‌سی", Price: 4200000, Active: true},
	{ID: "i3", Code: "S-003", Name: "جوانسازی پوست", Kind: "service", Unit: "جلسه", Price: 5500000, Active: true},
	{ID: "i4", Code: "S-004", Name: "مشاوره تخصصی", Kind: "service", Unit: "جلسه", Price: 800000, Active: true},
	{ID: "i5", Code: "G-001", Name: "سرم هیالورونیک", Kind: "goods", Unit: "عدد", Price: 1200000, StockQty: 24, Active: true},
	{ID: "i6", Code: "G-002", Name: "کرم ترمیم‌کننده", Kind: "goods", Unit: "عدد", Price: 650000, StockQty: 40, Active: true},
	{ID: "i7", Code: "G-003", Name: "ماسک کلاژن", Kind: "goods", Unit: "بسته", Price: 380000, StockQty: 15, Active: true},
}

func (m *Module) list(c *fiber.Ctx) error {
	kind := c.Query("kind") // service | goods | empty=all
	q := strings.TrimSpace(strings.ToLower(c.Query("q")))

	result := make([]Item, 0, len(demoItems))
	for _, item := range demoItems {
		if kind != "" && item.Kind != kind {
			continue
		}
		if q != "" &&
			!strings.Contains(strings.ToLower(item.Name), q) &&
			!strings.Contains(strings.ToLower(item.Code), q) {
			continue
		}
		result = append(result, item)
	}
	return c.JSON(fiber.Map{
		"data":  result,
		"total": len(result),
	})
}

func (m *Module) get(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, item := range demoItems {
		if item.ID == id {
			summary := ItemSummary{Item: item}
			if item.Kind == "goods" {
				summary.LastBuyPrice = item.Price * 70 / 100 // demo
				summary.LastSalePrice = item.Price
				summary.LastMovement = "2026-10-05"
			}
			return c.JSON(summary)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "item not found"})
}
