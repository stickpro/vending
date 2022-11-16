package main

import "github.com/stickpro/vending/internal/app"

const configDir = "configs"

func main() {
	app.Run(configDir)
}
