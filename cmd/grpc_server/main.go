package main

import (
	"context"
	"log"

	"github.com/HpPpL/microservices_course_auth/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err := app.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
