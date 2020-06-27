package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"log"
)

type EmailSenderImpl struct {
	sesRepository SesRepository
}

func InitEmailSender(sesRepository SesRepository) EmailSender {
	return EmailSenderImpl{sesRepository: sesRepository}
}

func (e EmailSenderImpl) SendEmail(payload model.TimeTrackingEmailPayload) {
	log.Printf("Sending email to %v with subject %v", payload.Recipient, payload.Subject)
	e.sesRepository.SendEmail(
		payload.Sender,
		payload.Recipient,
		payload.HtmlBody,
		payload.Subject)
}
