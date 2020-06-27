package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"github.com/ivanmartos/bamboo-tracker/internal/repository"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
	"time"
)

type TodaysTimesheetEntries struct {
	IsEmpty          bool                   `json:"isEmpty"`
	TimesheetEntries []model.TimesheetEntry `json:"timesheetEntries"`
}

func Handler(_ context.Context) (TodaysTimesheetEntries, error) {
	timesheetParser := timesheetUploader.InitTimesheetParser(repository.InitS3RepositoryImpl(repository.GetS3Client()))
	timesheetEntries := timesheetParser.GetTimesheetEntries(time.Now().Weekday())

	return TodaysTimesheetEntries{
		IsEmpty:          len(timesheetEntries) == 0,
		TimesheetEntries: timesheetEntries,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
