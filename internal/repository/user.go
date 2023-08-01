package repository

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"gorm.io/gorm"
)

type UserRepo struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewUserRepo(db *gorm.DB, logger logger.Logger) *UserRepo {
	return &UserRepo{db: db, logger: logger}
}

func (us *UserRepo) SaveUser(user entity.User) (entity.User, error) {
	log := us.logger.Named("Create").With("user", user)
	err := us.db.Transaction(func(tx *gorm.DB) error {
		var exists bool
		err := tx.Model(&entity.User{}).Select("count(*) > 0").Where("email = ?", user.Email).Find(&exists).Error
		if err != nil {
			return err
		}
		if exists {
			return errs.NewError(errs.AlreadyExist, "user with such mail already exists")
		}
		return tx.Create(&user).Error
	})
	log.Debug("successfully created user")
	return user, err
}

func (us *UserRepo) GetUserByEmailAndPassword(email, password string) (entity.User, error) {
	log := us.logger.Named("GetUserByEmailAndPassword").With("email", email, "password", password)
	var user entity.User
	err := us.db.Where("email = ? AND password_hash = ?", email, password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errs.NewError(errs.NotFound, "User not found")
	}
	if err != nil {
		return user, err
	}
	log.Debug("successfully found user")
	return user, nil
}

func (us *UserRepo) GetUserById(id uuid.UUID) (entity.User, error) {
	log := us.logger.Named("GetUserById").With("id", id)
	var user entity.User
	err := us.db.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errs.NewError(errs.NotFound, "User not found")
	}
	if err != nil {
		return user, err
	}
	log.Debug("successfully found user")
	return user, nil
}
