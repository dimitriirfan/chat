package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/dimitriirfan/chat-2/internal/config"
)

type RESTConfig struct {
	Port int `env:"REST_PORT" envDefault:"8080"`
}

func main() {
	cfg := RESTConfig{}
	env.Parse(&cfg)
	r := config.NewRESTRouter()
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	httpServer.ListenAndServe()

}
