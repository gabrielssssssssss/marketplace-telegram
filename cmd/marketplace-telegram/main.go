package main

import (
	"os"

	_ "github.com/gabrielssssssssss/marketplace-telegram/docs"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

// @title           Swagger Marketplace API
// @version         1.0
// @description     This is a API documentation for use marketplace database.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:6060
// @BasePath  /api/v1
func main() {
	gotenv.Load(".env")
	controller.Controller()
}
