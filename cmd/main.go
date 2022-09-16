package main

import (
	// _ "github.com/bmizerany/pq"
	"github.com/joho/godotenv"
	meme "github.com/mihailco/memessenger"
	handler "github.com/mihailco/memessenger/pkg/handlers"
	"github.com/mihailco/memessenger/pkg/repository"
	"github.com/mihailco/memessenger/pkg/service"
	"github.com/sirupsen/logrus"

	// "github.com/sirupsen/logrusrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error in loading env: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	//'postgres://postgres:qwerty@localhost:5436?sslmode=disable'  up
	//postgres://postgres:postgrespw@localhost:49156
	if err != nil {
		logrus.Fatal("falled to init db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(meme.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
