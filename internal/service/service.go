package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
)

type Service struct {
	Authorization
	User
	Wishlist
	WishlistItem
}

type Authorization interface {
	CreateUser(name, email, password string) (entity.User, error)
	Authenticate(email string, password string) (string, error)
	ParseToken(token string) (entity.User, error)
}

type User interface {
}

type Wishlist interface {
	GetAllByUserID(id uuid.UUID) ([]entity.Wishlist, error)
	GetByID(id uuid.UUID) (entity.Wishlist, error)
	GetItemsByID(id uuid.UUID) ([]entity.WishlistItem, error)
	CreateWishlist(userId uuid.UUID, title, description string) (entity.Wishlist, error)
	UpdateWishlist(id uuid.UUID, title, description string) (entity.Wishlist, error)
	DeleteWishlist(id uuid.UUID) (entity.Wishlist, error)
}

type WishlistItem interface {
	GetByID(id uuid.UUID) (entity.WishlistItem, error)
	AddItemToWishlist(item entity.WishlistItem) (entity.WishlistItem, error)
	UpdateItem(id uuid.UUID, item entity.WishlistItem) (entity.WishlistItem, error)
	DeleteItem(id uuid.UUID) (entity.WishlistItem, error)
}

func NewService(repos *repository.Repository, logger logger.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repos.User, logger),
		Wishlist:      NewWishlistService(repos.Wishlist, logger),
		WishlistItem:  NewWishlistItemService(repos.WishlistItem, logger),
	}
}
