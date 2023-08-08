package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
)

type WishlistService struct {
	repo repository.Wishlist
	log  logger.Logger
}

func NewWishlistService(repo repository.Wishlist, log logger.Logger) *WishlistService {
	return &WishlistService{repo: repo, log: log}
}

func (ws *WishlistService) GetAllByUserID(userId uuid.UUID) ([]entity.Wishlist, error) {
	return ws.repo.GetAllByUserID(userId)
}

func (ws *WishlistService) GetByID(id uuid.UUID) (entity.Wishlist, error) {
	return ws.repo.GetByID(id)
}

func (ws *WishlistService) GetItemsByID(id uuid.UUID) ([]entity.WishlistItem, error) {
	return ws.repo.GetItemsByID(id)
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
	return ws.repo.UpdateWishlist(id, entity.Wishlist{
		Base: entity.Base{
			ID: id,
		},
		Title:       title,
		Description: description,
	})
}

func (ws *WishlistService) DeleteWishlist(id uuid.UUID) (entity.Wishlist, error) {
	return ws.repo.DeleteWishlist(id)
}
