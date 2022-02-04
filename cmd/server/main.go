package main

import "fmt"

// App - the struct which contains things like pointer
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up Our App")
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
