package main

import (
	"fmt"
	"net/http"

	transportHttp "production-ready-api/internal/transport/http"
)

// App - the struct which contains things like pointer
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up Our App")

	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8000", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go RestAPI")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our RESTAPI")
		fmt.Println(err)
	}
}
