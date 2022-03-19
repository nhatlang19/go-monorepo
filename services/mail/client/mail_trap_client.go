package client

import (
	"log"
	"net/smtp"
)

type MailTrap interface {
	Send(emailTo string, content []byte) (bool, error)
}

type mailTrap struct {
	Username  string
	Password  string
	EmailFrom string
}

func NewMailTrap(username string, password string, emailFrom string) MailTrap {
	return mailTrap{
		Username:  username,
		Password:  password,
		EmailFrom: emailFrom,
	}
}

func (m mailTrap) Send(emailTo string, content []byte) (bool, error) {
	auth := smtp.PlainAuth("", m.Username, m.Password, "smtp.mailtrap.io")
	log.Printf("%v", m)
	to := []string{emailTo}
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, m.EmailFrom, to, content)
	if err != nil {
		return false, err
	}

	return true, nil
}
