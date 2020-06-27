package timesheetUploader

import (
	"fmt"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"time"
)

type EmailComposerImpl struct {
}

func (e EmailComposerImpl) ComposeTimeTrackingEmailPayload(timeTracking model.TimeTracking) model.TimeTrackingEmailPayload {
	subject := fmt.Sprintf("BambooHr timesheet report for %s %s", time.Now().Weekday().String(), time.Now().Format("02.01.2006"))

	if timeTracking.DailyTotal == 0 {
		subject = "Empty " + subject
	}

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

	return model.TimeTrackingEmailPayload{
		Sender:    getEnvVariable("DAILY_TIME_TRACKING_SENDER_EMAIL"),
		Recipient: getEnvVariable("DAILY_TIME_TRACKING_RECIPIENT_EMAIL"),
		HtmlBody:  html,
		Subject:   subject,
	}
}
