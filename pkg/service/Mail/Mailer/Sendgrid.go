package Mailer

import (
  MailMessage "p2/pkg/service/Mail/Message"

  "fmt"
  "log"

  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridMailer struct {
  apiKey string
  username string
}

func NewSendgrid(username, apiKey string) *sendgridMailer {
  return &sendgridMailer{apiKey: apiKey, username: username}
}

func (m *sendgridMailer) Send(message MailMessage.Message) {
  from := mail.NewEmail(m.username, m.username)
  subject := "REMINDER"
  to := mail.NewEmail(message.GetTo()[0], message.GetTo()[0])
  msgMail := mail.NewSingleEmail(from, subject, to, message.GetBody(), message.GetBody())
  client := sendgrid.NewSendClient(m.apiKey)
  response, err := client.Send(msgMail)
  if err != nil {
    log.Println("sendgridMailer send error", err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}
