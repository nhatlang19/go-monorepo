package service

import (
	"github.com/nhatlang19/go-monorepo/services/mail/client"
)

type RegisterService interface {
	Send(emailTo string) (bool, error)
}

type registerService struct {
	MailTrap client.MailTrap
}

func NewRegisterService(m client.MailTrap) RegisterService {
	return registerService{
		MailTrap: m,
	}
}

func (r registerService) Send(emailTo string) (bool, error) {
	msg := []byte("To: " + emailTo + "\r\n" +
		"Subject: Register successful?\r\n" +
		"\r\n" +
		"Here’s the space for our great sales pitch\r\n")
	return r.MailTrap.Send(emailTo, msg)
}
