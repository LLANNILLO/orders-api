package main

import (
	"context"
	"fmt"

	"github.com/llannillo/orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.Background())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
