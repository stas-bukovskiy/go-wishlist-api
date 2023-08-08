package service

import (
	uuid "github.com/satori/go.uuid"

	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/uploader"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"mime/multipart"
)

type ImageService struct {
	uploader uploader.Image
	repo     repository.Image
	log      logger.Logger
}

func NewImageService(uploader uploader.Image, repo repository.Image, log logger.Logger) *ImageService {
	return &ImageService{uploader: uploader, repo: repo, log: log}
}

func (s *ImageService) CreateImage(file multipart.File, header *multipart.FileHeader) (entity.Image, error) {
	imageID := uuid.NewV4()
	url, imageName, err := s.uploader.UploadImage(file, header, imageID)
	if err != nil {
		return entity.Image{}, err
	}

	return s.repo.SaveImage(entity.Image{
		Base:           entity.Base{ID: imageID},
		URL:            url,
		ImageName:      imageName,
		WishlistItemID: uuid.Nil,
	})
}

func (s *ImageService) DeleteImage(id uuid.UUID) error {
	image, err := s.repo.GetImage(id)
	if err != nil {
		return err
	}
	err = s.uploader.DeleteImage(image.ImageName)
	if err != nil {
		return err
	}
	return s.repo.DeleteImage(id)
}
