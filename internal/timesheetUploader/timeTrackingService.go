package timesheetUploader

import (
	"fmt"
	"log"
	"time"
)

type TimeTrackingService struct {
	api           BambooApi
	sesRepository SesRepository
}

func InitTimeTrackingService(api BambooApi, sesRepository SesRepository) TimeTrackingService {
	return TimeTrackingService{
		api:           api,
		sesRepository: sesRepository,
	}
}

func (s TimeTrackingService) SendCurrentTimeTrackEmail() {
	s.api.LogIn(getEnvVariable("BAMBOO_USERNAME"), getEnvVariable("BAMBOO_PASSWORD"))
	timeTracking := s.api.GetHomeContent()

	subject := fmt.Sprintf("BambooHr timesheet report for %s %s", time.Now().Weekday().String(), time.Now().Format("02.01.2006"))

	htmlTemplate := `
<html>
<body>
<h3>%s</h3>
<p>Logged time for today - %f</p>
<p>Logged time for the week - %f</p>
</body>
</html>
`
	html := fmt.Sprintf(htmlTemplate, subject, timeTracking.DailyTotal, timeTracking.WeeklyTotal)

	log.Println(html)

	s.sesRepository.SendEmail(
		getEnvVariable("DAILY_TIME_TRACKING_SENDER_EMAIL"),
		getEnvVariable("DAILY_TIME_TRACKING_RECIPIENT_EMAIL"),
		html,
		subject)

	log.Println(timeTracking.DailyTotal)
}
