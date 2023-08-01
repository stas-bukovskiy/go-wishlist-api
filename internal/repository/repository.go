package repository

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"gorm.io/gorm"
)

type Repository struct {
	User
}

func NewRepository(db *gorm.DB, logger logger.Logger) *Repository {
	return &Repository{
		User: NewUserRepo(db, logger),
	}
}

type User interface {
	SaveUser(user entity.User) (entity.User, error)
	GetUserByEmailAndPassword(email string, password string) (entity.User, error)
	GetUserById(id uuid.UUID) (entity.User, error)
}
