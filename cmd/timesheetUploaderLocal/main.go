package main

import (
	"github.com/ivanmartos/bamboo-tracker/internal/repository"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

func main() {
	timesheetService := timesheetUploader.InitTimesheetService(
		repository.InitBambooApi(),
		timesheetUploader.InitTimesheetParser(repository.InitS3RepositoryImpl(repository.GetS3Client())),
	)

	timesheetService.UploadCurrentTimesheet()
}
