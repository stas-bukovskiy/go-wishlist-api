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
	err := w.db.Preload("Images").First(&item, id).Error
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

		imageIDs := fromImagesToIds(item.Images)
		err = tx.Model(&entity.Image{}).Select("count(*) = ?", len(item.Images)).Where("id IN (?)", imageIDs).Find(&exist).Error
		if err != nil {
			return err
		}
		if !exist {
			return errs.NewError(errs.NotFound, "some image is not found")
		}

		err = tx.Omit("Images").Save(&item).Error
		if err != nil {
			return err
		}

		err = tx.Model(&entity.Image{}).Where("id IN (?)", imageIDs).Update("WishlistItemID", item.ID).Error
		if err != nil {
			return err
		}

		return tx.Preload("Images").First(&item, item.ID).Error
	})
	return item, err
}

func fromImagesToIds(images []entity.Image) []uuid.UUID {
	var ids []uuid.UUID
	for _, image := range images {
		ids = append(ids, image.ID)
	}
	return ids
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

		imageIDs := fromImagesToIds(item.Images)
		err = tx.Model(&entity.Image{}).Select("count(*) = ?", len(item.Images)).Where("id IN (?)", imageIDs).Find(&exist).Error
		if err != nil {
			return err
		}
		if !exist {
			return errs.NewError(errs.NotFound, "some image is not found")
		}

		err = tx.Model(&entity.Image{}).Where("wishlist_item_id = ?", id).Updates(map[string]interface{}{
			"wishlist_item_id": nil,
		}).Error
		if err != nil {
			return err
		}
		err = tx.Model(&entity.Image{}).Where("id IN (?)", imageIDs).Update("WishlistItemID", id).Error
		if err != nil {
			return err
		}

		err = tx.Omit("Images").Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
			"title":       item.Title,
			"description": item.Description,
			"price":       item.Price,
		}).Error
		if err != nil {
			return err
		}
		return tx.Preload("Images").Find(&item, id).Error
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

		err = tx.Clauses(clause.Returning{}).Preload("Images").Where("id = ?", id).Delete(&item).Error
		if err != nil {
			return err
		}
		return tx.Preload("Images").Find(&item, id).Error
	})
	return item, err
}
