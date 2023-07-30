package service

import (
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
)

type Service struct {
	repos  *repository.Repository
	logger logger.Logger
}

func NewService(repos *repository.Repository, logger logger.Logger) *Service {
	return &Service{repos: repos, logger: logger}
}
