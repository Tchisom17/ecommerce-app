package domain

type ItemStatus bool

const (
	Available    ItemStatus = true
	NotAvailable ItemStatus = false
)

type Item struct {
	Models
	UserID      string `json:"user_id"`
	User        User
	Title       string `json:"title" gorm:"not null"`
	CategoryID  string `json:"category_id"`
	Category    Category
	Description string     `json:"description"  gorm:"not null"`
	Price       int        `json:"price"  gorm:"not null"`
	Location    string     `json:"location" gorm:"not null"`
	ItemStatus  ItemStatus `json:"item_status" gorm:"not null default:true"`
	Images      []Images   `json:"images" gorm:"not null"`
}

type BookmarkItem struct {
	UserID string `json:"user_id"`
	ItemID string `json:"item_id"`
}
