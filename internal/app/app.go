package app

import (
	"fiolib/internal/config"
	"fiolib/internal/container"
	"fiolib/internal/logger"
	connector "fiolib/storage"
	"fiolib/storage/migrations"
	"log"

	server "fiolib/internal/transport/http"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var cfg *config.Config

func init() {
	cfg = config.MustLoad()
}

func Start() {
	apilogger := logger.GetAPILogger(cfg.Logger.APIlp)
	dblogger := logger.GetDBLogger(cfg.Logger.DBlp)

	db := checkUpDB(dblogger)

	container := container.NewContainer(db)

	server := server.NewHTTPServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port, container, apilogger)

	apilogger.Printf("Set IP: %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	apilogger.Println("Starting server")

	log.Printf("Server work on %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func checkUpDB(logger gormlogger.Interface) *gorm.DB {
	db := connector.Psql(cfg.Db.Role, cfg.Db.Pass, cfg.Db.Name, cfg.Db.Host, cfg.Db.Port, logger)
	migrations.RunMigration(db)

	return db
}
