package entity

import uuid "github.com/satori/go.uuid"

type Wishlist struct {
	Base
	Title       string         `json:"title"`
	Description string         `gorm:"size:2000" json:"description"`
	UserID      uuid.UUID      `gorm:"<-:create" json:"user_id"`
	User        User           `json:"-"`
	Items       []WishlistItem `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}

type WishlistItem struct {
	Base
	WishlistId  uuid.UUID `json:"wishlist_id"`
	Title       string    `json:"title"`
	Description string    `gorm:"size:2000" json:"description"`
	Price       uint      `json:"price"`
	Images      []Image   `json:"images" gorm:"constraint:OnDelete:CASCADE"`
	URL         string    `json:"url"`
}

type Image struct {
	Base
	URL            string    `json:"url"`
	ImageName      string    `json:"image_name"`
	WishlistItemID uuid.UUID `json:"wishlist_item_id"`
}
