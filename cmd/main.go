package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/handler"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/service"
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
	err = db.DB.AutoMigrate()
	if err != nil {
		log.With("error", err).Fatal("error occurred during database schemas migration")
	}

	// layers initialization
	repos := repository.NewRepository(db.DB, log)
	services := service.NewService(repos, log)
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
