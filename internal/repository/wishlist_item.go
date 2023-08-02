package repository

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WishlistItemRepo struct {
	db  *gorm.DB
	log logger.Logger
}

func NewWishlistItemRepo(db *gorm.DB, log logger.Logger) *WishlistItemRepo {
	return &WishlistItemRepo{db: db, log: log}
}

func (w *WishlistItemRepo) GetByID(id uuid.UUID) (entity.WishlistItem, error) {
	var item entity.WishlistItem
	err := w.db.First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.WishlistItem{}, errs.NewError(errs.NotFound, "wishlist item wish such id not found")
		}
		return entity.WishlistItem{}, err
	}
	return item, nil
}

func (w *WishlistItemRepo) CreateWishlistItem(item entity.WishlistItem) (entity.WishlistItem, error) {
	err := w.db.Transaction(func(tx *gorm.DB) error {
		var exist bool
		err := tx.Model(&entity.Wishlist{}).Select("count(*) > 0").Where("id = ?", item.WishlistId).Find(&exist).Error
		if err != nil {
			return err
		}
		if !exist {
			return errs.NewError(errs.NotFound, "wishlist with such id not found")
		}

		return tx.Create(&item).Error
	})
	return item, err
}

func (w *WishlistItemRepo) UpdateItem(id uuid.UUID, item entity.WishlistItem) (entity.WishlistItem, error) {
	err := w.db.Transaction(func(tx *gorm.DB) error {
		var exist bool
		err := tx.Model(&entity.WishlistItem{}).Select("count(*) > 0").Where("id = ?", id).Find(&exist).Error
		if err != nil {
			return err
		}
		if !exist {
			return errs.NewError(errs.NotFound, "wishlist item with such id not found")
		}

		return tx.Model(&item).Clauses(clause.Returning{}).Where("id = ?", id).Updates(map[string]interface{}{
			"title":       item.Title,
			"description": item.Description,
			"price":       item.Price,
			//"image_urls":  item.ImageURLs,
		}).Error
	})
	return item, err
}

func (w *WishlistItemRepo) DeleteItem(id uuid.UUID) (entity.WishlistItem, error) {
	var item entity.WishlistItem
	err := w.db.Transaction(func(tx *gorm.DB) error {
		var exist bool
		err := tx.Model(&entity.WishlistItem{}).Select("count(*) > 0").Where("id = ?", id).Find(&exist).Error
		if err != nil {
			return err
		}
		if !exist {
			return errs.NewError(errs.NotFound, "wishlist item with such id not found")
		}

		return tx.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&item).Error
	})
	return item, err
}
