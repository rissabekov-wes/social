package main

import (
	"github.com/Wesfarmers-Digital/pkg/one_http"
	"github.com/rissabekov-wes/social/internal/api"
	"github.com/rissabekov-wes/social/internal/config"
)

func main() {
	appConfig := config.NewApplicationConfig()

	srv := one_http.NewServer(appConfig.ServiceName())
	srv.DisableTLS = true
	srv.Port = appConfig.ServerPort()
	srv.RegisterRoute(
		api.ConfigRoute(),
	)

	srv.Start()
}
