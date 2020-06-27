package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

func Handler(_ context.Context, timeTracking model.TimeTracking) (model.TimeTrackingEmailPayload, error) {
	emailComposer := timesheetUploader.EmailComposerImpl{}

	return emailComposer.ComposeTimeTrackingEmailPayload(timeTracking), nil
}

func main() {
	lambda.Start(Handler)
}
