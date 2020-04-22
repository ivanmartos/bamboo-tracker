package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ivanmartos/bamboo-tracker/internal/repository"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(_ context.Context) {
	timesheetService := timesheetUploader.InitTimesheetService(
		repository.InitBambooApi(),
		timesheetUploader.InitTimesheetParser(),
	)

	timesheetService.UploadCurrentTimesheet()
}

func main() {
	lambda.Start(Handler)
}
