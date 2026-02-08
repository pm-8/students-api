package config

import (
	"flag"
	"os"
	"log"
)

type Http_Server struct {
	Addr string
}

// env-default : "production"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"` //struct tags
	StoragePath string `yaml:"storage_path" env-required:"true"`
	Http_Server `yaml:"http_server"`
}

func MustLoad() { //things that should load to make this api work
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config_path","","path to config file")
		flag.Parse()
		configPath = *flags
		if configPath == ""{
			log.Fatal("Config Path is not set")
		}
	}
}