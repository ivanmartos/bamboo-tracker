package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"testing"
)

func TestEmailSenderImpl_SendEmail(t *testing.T) {
	type fields struct {
		sesRepository mocks.MockSesRepository
	}
	type args struct {
		payload model.TimeTrackingEmailPayload
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Sending email",
			fields: fields{
				sesRepository: mocks.MockSesRepository{},
			},
			args: args{
				payload: model.TimeTrackingEmailPayload{
					Sender:    "sender@mail.com",
					Recipient: "recipient@email.com",
					HtmlBody:  "html body",
					Subject:   "anySubject",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EmailSenderImpl{
				sesRepository: &tt.fields.sesRepository,
			}
			e.SendEmail(tt.args.payload)
			if tt.fields.sesRepository.SendEmailHtmlBodyArg != tt.args.payload.HtmlBody {
				t.Errorf("SendEmail got HTML %v, want %v", tt.fields.sesRepository.SendEmailHtmlBodyArg, tt.args.payload.HtmlBody)
			}

			if tt.fields.sesRepository.SendEmailRecipientArg != tt.args.payload.Recipient {
				t.Errorf("SendEmail got recipient %v, want %v", tt.fields.sesRepository.SendEmailRecipientArg, tt.args.payload.Recipient)
			}

			if tt.fields.sesRepository.SendEmailSenderArg != tt.args.payload.Sender {
				t.Errorf("SendEmail got sender %v, want %v", tt.fields.sesRepository.SendEmailSenderArg, tt.args.payload.Sender)
			}

			if tt.fields.sesRepository.SendEmailSubjectArg != tt.args.payload.Subject {
				t.Errorf("SendEmail got sender %v, want %v", tt.fields.sesRepository.SendEmailSubjectArg, tt.args.payload.Subject)
			}
		})
	}
}
