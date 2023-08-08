package request

import uuid "github.com/satori/go.uuid"

type SignUpRequest struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,min=4,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,min=4,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AddWishlistItemRequest struct {
	WishlistId  uuid.UUID   `json:"wishlist_id" binding:"required"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       uint        `json:"price"`
	ImageIDs    []uuid.UUID `json:"image_ids" binding:"max=5"`
	URL         string      `json:"url" binding:"required,url"`
}

type UpdateWishlistItemRequest struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       uint        `json:"price"`
	ImageIDs    []uuid.UUID `json:"image_ids" binding:"max=5"`
}

type WishlistRequest struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description"`
}
