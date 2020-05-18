package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ivanmartos/bamboo-tracker/internal/repository"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(_ context.Context) {
	timeTrackingService := timesheetUploader.InitTimeTrackingService(
		repository.InitBambooApi(),
		repository.InitSesRepositoryImpl(repository.GetSesClient()),
	)

	timeTrackingService.SendCurrentTimeTrackEmail()
}

func main() {
	lambda.Start(Handler)
}
