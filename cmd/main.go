package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/Futturi/GolangSchoolProject/internal/handler"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
	"github.com/Futturi/GolangSchoolProject/internal/server"
	"github.com/Futturi/GolangSchoolProject/internal/service"
	"github.com/Futturi/GolangSchoolProject/pkg"
	repositoryinitf "github.com/Futturi/GolangSchoolProject/pkg/repository_initf"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("error while reading .env file", err)
	}
	if err := InitConfig(); err != nil {
		slog.Error("error while reading config: ", err)
	}
	cfg := repositoryinitf.Config{
		Hostname: viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		NameDB:   viper.GetString("db.namedb"),
		Sslmode:  viper.GetString("db.sslmode"),
	}
	rcfg := pkg.RedisConf{
		Addr:     viper.GetString("red.host") + viper.GetString("red.port"),
		Password: "",
		Db:       0,
	}
	rdb := pkg.NewRedis(rcfg)
	db, err := repositoryinitf.InitPostgre(cfg)
	if err != nil {
		slog.Error("error while connecting to db", err)
	}
	repo := repository.NewReposiotry(db, rdb)
	serv := service.NewService(repo)
	hand := handler.NewHandler(serv)

	serve := new(server.Server)
	if err = serve.InitServer(viper.GetString("port"), hand.InitRoutes()); err != nil {
		log.Fatal("erorr")
	}
}

func InitConfig() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("internal/configs")
	return viper.ReadInConfig()
}
