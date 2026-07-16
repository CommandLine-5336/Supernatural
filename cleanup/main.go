package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	application "github.com/CommandLine-5336/Supernatural/cleanup/app"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Printf("Error starting the application: %v\n", err)
	}
}
