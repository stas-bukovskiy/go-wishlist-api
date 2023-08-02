package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
)

type WishlistItemService struct {
	repo   repository.WishlistItem
	logger logger.Logger
}

func NewWishlistItemService(repo repository.WishlistItem, logger logger.Logger) *WishlistItemService {
	return &WishlistItemService{repo: repo, logger: logger}
}

func (w *WishlistItemService) GetByID(id uuid.UUID) (entity.WishlistItem, error) {
	item, err := w.repo.GetByID(id)
	return item, err
}

func (w *WishlistItemService) AddItemToWishlist(item entity.WishlistItem) (entity.WishlistItem, error) {
	item, err := w.repo.CreateWishlistItem(item)
	return item, err
}

func (w *WishlistItemService) UpdateItem(id uuid.UUID, item entity.WishlistItem) (entity.WishlistItem, error) {
	item, err := w.repo.UpdateItem(id, item)
	return item, err
}

func (w *WishlistItemService) DeleteItem(id uuid.UUID) (entity.WishlistItem, error) {
	item, err := w.repo.DeleteItem(id)
	return item, err
}
