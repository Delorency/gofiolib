package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ConfigHTTPServer struct {
	Host string `env:"HOST" env-default:"localhost"`
	Port string `env:"PORT" env-default:"8080"`
}

type ConfigDatabase struct {
	Role string `env:"DB_ROLE"`
	Pass string `env:"DB_PASS"`
	Name string `env:"DB_NAME"`
	Host string `env:"DB_HOST"`
	Port string `env:"DB_PORT"`
}

type ConfigLogger struct {
	APIlp   string `env:"APILOGFILENAME"`
	DBlp    string `env:"DBLOGFILENAME"`
	LogsDir string `env:"LOGSDIR"`
}

type Config struct {
	HTTPServer *ConfigHTTPServer
	Db         *ConfigDatabase
	Logger     *ConfigLogger
}

func MustLoad() *Config {
	godotenv.Load("./configs/.env")

	var cfgHttpServer ConfigHTTPServer
	var cfgDatabase ConfigDatabase
	var cgfLogger ConfigLogger

	if err := cleanenv.ReadEnv(&cfgHttpServer); err != nil {
		log.Fatalln("Ошибка чтения настроек сервера из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cfgDatabase); err != nil {
		log.Fatalln("Ошибка чтения настроек бд из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cgfLogger); err != nil {
		log.Fatalln("Ошибка чтения настроек логгера из .env файлы")
	}

	return &Config{&cfgHttpServer, &cfgDatabase, &cgfLogger}
}
