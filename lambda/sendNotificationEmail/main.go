package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"github.com/ivanmartos/bamboo-tracker/internal/repository"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

func Handler(_ context.Context, emailPayload model.TimeTrackingEmailPayload) {
	service := timesheetUploader.InitEmailSender(repository.InitSesRepositoryImpl(repository.GetSesClient()))

	service.SendEmail(emailPayload)
}

func main() {
	lambda.Start(Handler)
}
