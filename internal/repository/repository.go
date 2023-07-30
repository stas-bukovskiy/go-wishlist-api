package repository

import (
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"gorm.io/gorm"
)

type Repository struct {
	User
}

func NewRepository(db *gorm.DB, logger logger.Logger) *Repository {
	return &Repository{}
}

type User interface {
}
