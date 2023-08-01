package entity

import uuid "github.com/satori/go.uuid"

type WishlistItem struct {
	Base
	WishlistId  uuid.UUID `json:"wishlist_id"`
	Title       string    `json:"title"`
	Description string    `gorm:"size:2000" json:"description"`
	Price       uint      `json:"price"`
	URL         string    `json:"url"`
}

type Wishlist struct {
	Base
	Title       string         `json:"title"`
	Description string         `gorm:"size:2000" json:"description"`
	UserID      uuid.UUID      `gorm:"<-:create" json:"user_id"`
	User        User           `json:"-"`
	Items       []WishlistItem `json:"items,omitempty"`
}
