package uploader

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"net/url"

	"github.com/minio/minio-go/v7"

	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type Image interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, imageID uuid.UUID) (string, string, error)
	DeleteImage(imageName string) error
}

type ImageUploader struct {
	minioClient *minio.Client
	log         logger.Logger
}

func NewImageUploader(minioClient *minio.Client, logger logger.Logger) *ImageUploader {
	return &ImageUploader{minioClient: minioClient, log: logger}
}

func (u *ImageUploader) UploadImage(file multipart.File, header *multipart.FileHeader, imageID uuid.UUID) (string, string, error) {
	size, err := getFileSize(file)
	if err != nil {
		return "", "", err
	}

	extension := getExtension(header.Filename)
	fileName := imageID.String() + "." + extension

	info, err := u.minioClient.PutObject(context.Background(), "test", fileName, file, size, minio.PutObjectOptions{ContentType: "image/" + extension})
	u.log.Info("uploading image with size" + strconv.FormatInt(info.Size, 10))

	imageUrl, err := u.minioClient.PresignedGetObject(context.Background(), "test", fileName, time.Duration(1000)*time.Second, make(url.Values))
	return imageUrl.String(), fileName, err
}

func (u *ImageUploader) DeleteImage(imageName string) error {
	return u.minioClient.RemoveObject(context.Background(), "test", imageName, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
}

type Sizer interface {
	Size() int64
}

func getFileSize(file multipart.File) (int64, error) {
	fileHeader := make([]byte, 512)
	if _, err := file.Read(fileHeader); err != nil {
		return 0, err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return 0, err
	}
	return file.(Sizer).Size(), nil
}

func getExtension(fileName string) string {
	return strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])
}
