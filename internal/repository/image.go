package repository

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"gorm.io/gorm"
)

type ImageRepo struct {
	db  *gorm.DB
	log logger.Logger
}

func NewImageRepo(db *gorm.DB, log logger.Logger) *ImageRepo {
	return &ImageRepo{db: db, log: log}
}

func (r *ImageRepo) GetImage(id uuid.UUID) (entity.Image, error) {
	var image entity.Image
	err := r.db.First(&image, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Image{}, errs.NewError(errs.NotFound, "image wish such id not found")
		}
		return entity.Image{}, err
	}
	return image, nil
}

func (r *ImageRepo) SaveImage(image entity.Image) (entity.Image, error) {
	err := r.db.Omit("WishlistItemID").Save(&image).Error
	if err != nil {
		return entity.Image{}, err
	}
	return image, nil
}

func (r *ImageRepo) DeleteImage(id uuid.UUID) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var exist bool
		err := tx.Model(&entity.Image{}).Select("count(*) > 0").Where("id = ?", id).Find(&exist).Error
		if err != nil {
			return err
		}
		if !exist {
			return errs.NewError(errs.NotFound, "image with such id not found")
		}

		if err != nil {
			return err
		}
		return tx.Delete(&entity.Image{}, id).Error
	})
	return err
}
