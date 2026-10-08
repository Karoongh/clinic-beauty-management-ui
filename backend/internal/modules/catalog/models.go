package catalog

// Item represents a service or goods entry in the catalog.
type Item struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Kind      string `json:"kind"` // service | goods
	Unit      string `json:"unit"`
	Price     int64  `json:"price"` // default sell price in Rials
	StockQty  int    `json:"stock_qty,omitempty"` // only for goods
	Active    bool   `json:"active"`
}

// ItemSummary is the detail sheet shown when an item is selected.
type ItemSummary struct {
	Item
	LastBuyPrice  int64  `json:"last_buy_price,omitempty"`
	LastSalePrice int64  `json:"last_sale_price,omitempty"`
	LastMovement  string `json:"last_movement,omitempty"`
}
