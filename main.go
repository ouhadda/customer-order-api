package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ouhadda/customer-orders-api/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start the application:", err)
	}
}
