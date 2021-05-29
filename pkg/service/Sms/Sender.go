package Sms

import (
  SmsMessage "p2/pkg/service/Sms/Message"
)

type Sender interface {
  Send(message SmsMessage.Message)
}