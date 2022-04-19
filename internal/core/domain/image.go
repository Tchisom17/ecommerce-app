package domain

type Images struct {
	Models
	ItemID string `json:"item_id"`
	Item   Item
	URL    string `json:"url"`
	Name   string `json:"name"`
}
