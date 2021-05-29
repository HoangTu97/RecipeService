package Mail

import (
  "p2/pkg/service/Mail/Message"
)

type Mailer interface {
  Send(message MailMessage.Message)
}