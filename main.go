package main

import (
	"context"
	"fmt"

	"github.com/ouhadda/customer-orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start the application:", err)
	}
}
