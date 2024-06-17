package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/dimitriirfan/chat-2/internal/config"
)

type WebsocketConfig struct {
	Port int `env:"WS_PORT" envDefault:"443"`
}

func main() {
	cfg := WebsocketConfig{}
	env.Parse(&cfg)
	r := config.NewWebsocketRouter()
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	log.Printf("server started on port: %v", 8080)
	httpServer.ListenAndServe()

}
