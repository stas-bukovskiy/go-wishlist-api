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

type WishlistRepo struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewWishlistRepo(db *gorm.DB, logger logger.Logger) *WishlistRepo {
	return &WishlistRepo{db: db, logger: logger}
}

func (wr *WishlistRepo) GetAllByUserID(userId uuid.UUID) ([]entity.Wishlist, error) {
	var wishlists []entity.Wishlist
	err := wr.db.Where("user_id = ?", userId).Find(&wishlists).Error
	if err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (wr *WishlistRepo) GetByID(id uuid.UUID) (entity.Wishlist, error) {
	var wishlist entity.Wishlist
	err := wr.db.Where("id = ?", id).First(&wishlist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Wishlist{}, errs.NewError(errs.NotFound, "wishlist with such id not found")
		}
		return entity.Wishlist{}, err
	}
	return wishlist, nil
}

func (wr *WishlistRepo) GetItemsByID(id uuid.UUID) ([]entity.WishlistItem, error) {
	var items []entity.WishlistItem
	err := wr.db.Where("wishlist_id = ?", id).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (wr *WishlistRepo) CreateWishlist(wishlist entity.Wishlist) (entity.Wishlist, error) {
	err := wr.db.Save(&wishlist).Error
	if err != nil {
		return entity.Wishlist{}, err
	}
	return wishlist, nil
}

func (wr *WishlistRepo) UpdateWishlist(id uuid.UUID, wishlist entity.Wishlist) (entity.Wishlist, error) {
	err := wr.db.Transaction(func(tx *gorm.DB) error {
		var exists bool
		err := tx.Model(&entity.Wishlist{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
		if err != nil {
			return err
		}
		if !exists {
			return errs.NewError(errs.NotFound, "such wishlist does not exist")
		}
		wishlist.ID = id
		return tx.Model(&wishlist).Clauses(clause.Returning{}).Updates(map[string]interface{}{
			"title":       wishlist.Title,
			"description": wishlist.Description,
		}).Error
	})
	return wishlist, err
}

func (wr *WishlistRepo) DeleteWishlist(id uuid.UUID) (entity.Wishlist, error) {
	var wishlist entity.Wishlist
	err := wr.db.Transaction(func(tx *gorm.DB) error {
		var exists bool
		err := tx.Model(&entity.Wishlist{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
		if err != nil {
			return err
		}
		if !exists {
			return errs.NewError(errs.NotFound, "such wishlist does not exist")
		}
		return tx.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&wishlist).Error
	})
	return wishlist, err
}
