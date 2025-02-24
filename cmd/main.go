package main

import (
	"AnalyseService/analytics/handler"
	"AnalyseService/analytics/repository"
	"AnalyseService/analytics/services"
	"AnalyseService/cmd/server"
	"AnalyseService/pkg/database"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatal("error init config %v ", err)
	}

	srv := server.NewGRPCServer(viper.GetString("grpcSrv.port"))

	db, err := database.NewPostgresDB(database.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("cant run db %v", err)
	}
	defer db.Close()

	repo := repository.NewTaskRepo(db)
	service := services.NewAnalyticService(repo, repo)

	handler.Register(srv.GrpcServer, service)

	if err := srv.RunServer(); err != nil {
		log.Fatalf("cant run server %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
