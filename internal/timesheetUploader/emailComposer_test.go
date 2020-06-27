package timesheetUploader

import (
	"fmt"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestEmailComposerImpl_ComposeTimeTrackingEmailPayload(t *testing.T) {
	senderEmail := "foo@bar.lol"
	recipientEmail := "bar@foo.lol"

	_ = os.Setenv("DAILY_TIME_TRACKING_SENDER_EMAIL", senderEmail)
	_ = os.Setenv("DAILY_TIME_TRACKING_RECIPIENT_EMAIL", recipientEmail)

	type args struct {
		timeTracking model.TimeTracking
	}
	emptyHeader := fmt.Sprintf("Empty BambooHr timesheet report for %s %s", time.Now().Weekday().String(), time.Now().Format("02.01.2006"))
	filledHeader := fmt.Sprintf("BambooHr timesheet report for %s %s", time.Now().Weekday().String(), time.Now().Format("02.01.2006"))

	tests := []struct {
		name string
		args args
		want model.TimeTrackingEmailPayload
	}{
		{
			name: "Empty time tracking",
			args: args{model.TimeTracking{
				DailyTotal:  0,
				WeeklyTotal: 0,
			}},
			want: model.TimeTrackingEmailPayload{
				Sender:    senderEmail,
				Recipient: recipientEmail,
				Subject:   emptyHeader,
				HtmlBody: fmt.Sprintf(`
<html>
<body>
<h3>%v</h3>
<p>Logged time for today - 0.000000</p>
<p>Logged time for the week - 0.000000</p>
</body>
</html>
`, emptyHeader)}},
		{
			name: "Filled time tracking",
			args: args{model.TimeTracking{
				DailyTotal:  10,
				WeeklyTotal: 15,
			}},
			want: model.TimeTrackingEmailPayload{
				Sender:    senderEmail,
				Recipient: recipientEmail,
				Subject:   filledHeader,
				HtmlBody: fmt.Sprintf(`
<html>
<body>
<h3>%v</h3>
<p>Logged time for today - 10.000000</p>
<p>Logged time for the week - 15.000000</p>
</body>
</html>
`, filledHeader)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EmailComposerImpl{}
			if got := e.ComposeTimeTrackingEmailPayload(tt.args.timeTracking); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComposeTimeTrackingEmailPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}
