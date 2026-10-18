package main

import "fmt"

// to pass application configuration with optional values
// define a config of type struct
type AppConfig struct {
	Host string			// default value is ""
	Port int			// default value is 0
	UseTLS bool			// default value is false
	ConsoleMode bool
}

func startServer(cfg AppConfig) {
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	fmt.Println("App started using this config", cfg)
}

func main() {
	c := AppConfig{Host: "1.2.3.4", ConsoleMode: true}
	startServer(c)
}