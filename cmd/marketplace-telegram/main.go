package main

import (
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	gotenv.Load(".env")
	controller.Controller()
}
