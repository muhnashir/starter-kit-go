package provider

import (
	"starter-kit-go/config"
	"starter-kit-go/provider/database"
	"starter-kit-go/provider/p_routes"
	"time"
)

type App struct{}

func (app App) Start() {

	// Load environment
	LoadEnv()

	println("\n------------------------------------------------------------")
	println(config.App().Name + " app starting")
	println("------------------------------------------------------------\n")

	// Load database
	defer database.Close()
	database.Open()

	// Set time zone application
	time.LoadLocation(config.App().TimeZone)

	// Run route configuration
	p_routes.Run()
}
