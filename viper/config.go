package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

// This program uses Viper package to read a YAML configuration file, loads it into your Config struct
type Config struct {
	App struct {
		Name string
		Env	string
		Logging struct {
			LogName string
			LogLevel string
			LogMaxSize int64
		}
	}

	RabbitMQ struct {
		Host string
		Port int
		User string
		Pass string
		Exchange string
		ExchangeType string
		RoutingKey string
	}
}

func LoadConfig() (Config, error) {
	var appConfig Config
	if wd, err := os.Getwd(); err != nil {
		fmt.Println("Failed to get working directory")
	} else {
		fmt.Println("Current working dir: ", wd)
	}

	// configure Viper to read app.yml
	viper.SetConfigFile("viper/conf/app.yml")

	// read app.yml
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read config: %v", err)
		return appConfig, err
	}

	// load YAML values into the Config struct
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Printf("Failed to decode YAML config into struct: %v", err)
		return appConfig, err
	}

	// verify if the YAML is loaded properly
	fmt.Println("Env: ", appConfig.App.Env)
	fmt.Println("LogSize: ", appConfig.App.Logging.LogMaxSize)
	fmt.Println("RabbitMQ Exchange: ", appConfig.RabbitMQ.Exchange)

	return appConfig, nil
}