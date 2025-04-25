package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm/logger"
)

type APIL struct {
	Logger *log.Logger
	File   *os.File
}
type DBL struct {
	Logger logger.Interface
	File   *os.File
}

type CLogger struct {
	APIl *APIL
	DBl  *DBL
}

func MakeLoggerFile(path string) *os.File {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("Ошибка создания директории для логов: %v\n", err)
	}
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Ошибка создания лог-файла:", err)
	}

	return logFile
}
func GetLogger(apiFile, DbFile *os.File) *CLogger {
	logger := logger.New(
		log.New(DbFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
		},
	)
	loggerAPI := APIL{log.New(apiFile, "", log.LstdFlags), apiFile}
	loggerDB := DBL{logger, DbFile}

	return &CLogger{&loggerAPI, &loggerDB}
}
