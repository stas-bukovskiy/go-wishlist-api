package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/handler"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/service"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/uploader"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/database"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httpserver"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.New("DEBUG")

	// config initialization
	if err := initConfig(); err != nil {
		log.With("error", err).Fatal("error occurred during config initialization")
	}

	// .env file loading
	if err := godotenv.Load(); err != nil {
		log.With("error", err).Fatal("error occurred during .env file loading")
	}

	// database connection
	db, err := database.NewPostgreSQL(database.PostgreSQLConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetBool("db.sslmode"),
		TimeZone: viper.GetString("db.timezone"),
	})
	if err != nil {
		log.With("error", err).Fatal("error occurred during database connection")
	}

	// database schemas migration
	err = db.DB.AutoMigrate(entity.User{}, entity.Wishlist{}, entity.Image{}, entity.WishlistItem{})
	if err != nil {
		log.With("error", err).Fatal("error occurred during database schemas migration")
	}

	// image upload client initialization
	minioClient, err := minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_ID"), os.Getenv("MINIO_ACCESS_KEY"), ""),
		Secure: false,
	})
	if err != nil {
		log.With("error", err).Fatal("error occurred during minio client initialization")
	}

	exists, err := minioClient.BucketExists(context.Background(), "test")
	if err != nil {
		log.With("error", err).Fatal("error occurred during bucket existing check")
	}
	if !exists {
		err := minioClient.MakeBucket(context.Background(), "test", minio.MakeBucketOptions{})
		if err != nil {
			log.With("error", err).Fatal("error occurred during bucket creation")
		}
	}
	if err != nil {
		log.With("error", err).Fatal("error occurred during image upload client initialization")
	}

	// layers initialization
	repos := repository.NewRepository(db.DB, log)
	imageUploader := uploader.NewImageUploader(minioClient, log)
	services := service.NewService(repos, imageUploader, log)
	handlers := handler.NewHandler(services, log)

	// http server initialization
	srv := httpserver.New(handlers.InitRoutes(), httpserver.Port(viper.GetString("port")))
	log.Info("wishlist server has started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("wishlist server shutting down...")

	// server shutdown
	if err := srv.Shutdown(); err != nil {
		log.With("error", err).Error("error occurred while shutting down")
	}

	// db connection closing
	if err := db.Close(); err != nil {
		log.With("error", err).Error("error occurred while database connection closing")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
