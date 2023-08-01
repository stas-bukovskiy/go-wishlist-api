package service

import (
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
)

type Service struct {
	Authorization
	User
}

type Authorization interface {
	CreateUser(name, email, password string) (entity.User, error)
	Authorize(email string, password string) (string, error)
	ParseToken(token string) (entity.User, error)
}

type User interface {
}

func NewService(repos *repository.Repository, logger logger.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repos.User, logger),
	}
}
