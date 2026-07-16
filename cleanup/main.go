package main

import (
	"context"
	"fmt"

	application "github.com/CommandLine-5336/Supernatural/cleanup/app"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting the application: %v\n", err)
	}
}
