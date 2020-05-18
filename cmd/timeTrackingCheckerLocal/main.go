package main

import (
	"github.com/ivanmartos/bamboo-tracker/internal/repository"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

func main() {
	timeTrackingService := timesheetUploader.InitTimeTrackingService(
		repository.InitBambooApi(),
		repository.InitSesRepositoryImpl(repository.GetSesClient()),
	)

	timeTrackingService.SendCurrentTimeTrackEmail()
}
