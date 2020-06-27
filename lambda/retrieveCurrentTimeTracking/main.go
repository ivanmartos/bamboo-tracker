package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"github.com/ivanmartos/bamboo-tracker/internal/repository"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

func Handler(_ context.Context) (model.TimeTracking, error) {
	timeTrackingService := timesheetUploader.InitTimeTrackingService(repository.InitBambooApi())

	return timeTrackingService.GetCurrentTimeTracking(), nil
}

func main() {
	lambda.Start(Handler)
}
