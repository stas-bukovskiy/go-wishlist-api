package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"gorm.io/gorm"
)

type WishlistService struct {
	repo repository.Wishlist
	log  logger.Logger
}

func NewWishlistService(repo repository.Wishlist, log logger.Logger) *WishlistService {
	return &WishlistService{repo: repo, log: log}
}

func (ws *WishlistService) GetAllByUserID(userId uuid.UUID) ([]entity.Wishlist, error) {
	wishlists, err := ws.repo.GetAllByUserID(userId)
	if err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (ws *WishlistService) GetByID(id uuid.UUID) (entity.Wishlist, error) {
	wishlist, err := ws.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Wishlist{}, errs.NewError(errs.NotFound, "wishlist not found")
		}
		return entity.Wishlist{}, err
	}
	return wishlist, nil
}

func (ws *WishlistService) CreateWishlist(userId uuid.UUID, title, description string) (entity.Wishlist, error) {
	wishlist, err := ws.repo.CreateWishlist(entity.Wishlist{
		Title:       title,
		Description: description,
		UserID:      userId,
	})
	if err != nil {
		return entity.Wishlist{}, err
	}
	return wishlist, nil
}

func (ws *WishlistService) UpdateWishlist(id uuid.UUID, title, description string) (entity.Wishlist, error) {
	wishlist, err := ws.repo.UpdateWishlist(id, entity.Wishlist{
		Base: entity.Base{
			ID: id,
		},
		Title:       title,
		Description: description,
	})
	if err != nil {
		return entity.Wishlist{}, err
	}
	return wishlist, nil
}

func (ws *WishlistService) DeleteWishlist(id uuid.UUID) (entity.Wishlist, error) {
	wishlist, err := ws.repo.DeleteWishlist(id)
	if err != nil {
		return entity.Wishlist{}, err
	}
	return wishlist, nil
}
