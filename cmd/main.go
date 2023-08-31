package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/joho/godotenv"
	"github.com/naumovrus/backend-trainee-asignment/interntal/config"
	"github.com/naumovrus/backend-trainee-asignment/interntal/handler"
	"github.com/naumovrus/backend-trainee-asignment/interntal/httpserver"
	"github.com/naumovrus/backend-trainee-asignment/interntal/repository"
	"github.com/naumovrus/backend-trainee-asignment/interntal/service"
	"github.com/sirupsen/logrus"
)

func main() {

	// init config
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load .env files: %s", err.Error())
	}
	cfg := config.LoadConfig()

	db, err := repository.NewPostgresDB(repository.ConfigDB{
		Username: cfg.Username,
		Host:     cfg.Host,
		Port:     cfg.PortDb,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   cfg.Dbname,
		SSLMode:  cfg.Sslmode,
	})
	if err != nil {
		logrus.Fatalf("unabled to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(httpserver.Server)
	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("App Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s ", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
