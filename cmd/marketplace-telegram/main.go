package main

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(".env")
	controller.Controller()
}
