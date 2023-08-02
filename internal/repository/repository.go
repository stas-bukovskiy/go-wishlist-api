package repository

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"gorm.io/gorm"
)

type Repository struct {
	User
	Wishlist
	WishlistItem
}

func NewRepository(db *gorm.DB, logger logger.Logger) *Repository {
	return &Repository{
		User:         NewUserRepo(db, logger),
		Wishlist:     NewWishlistRepo(db, logger),
		WishlistItem: NewWishlistItemRepo(db, logger),
	}
}

type User interface {
	SaveUser(user entity.User) (entity.User, error)
	GetUserByEmailAndPassword(email string, password string) (entity.User, error)
	GetUserById(id uuid.UUID) (entity.User, error)
}

type Wishlist interface {
	GetAllByUserID(id uuid.UUID) ([]entity.Wishlist, error)
	GetByID(id uuid.UUID) (entity.Wishlist, error)
	CreateWishlist(wishlist entity.Wishlist) (entity.Wishlist, error)
	UpdateWishlist(id uuid.UUID, wishlist entity.Wishlist) (entity.Wishlist, error)
	DeleteWishlist(id uuid.UUID) (entity.Wishlist, error)
}

type WishlistItem interface {
	GetByID(id uuid.UUID) (entity.WishlistItem, error)
	CreateWishlistItem(item entity.WishlistItem) (entity.WishlistItem, error)
	UpdateItem(id uuid.UUID, item entity.WishlistItem) (entity.WishlistItem, error)
	DeleteItem(id uuid.UUID) (entity.WishlistItem, error)
}
