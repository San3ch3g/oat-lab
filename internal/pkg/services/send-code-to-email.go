package services

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"oat-lab-module/internal/utils/config"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, bcc []string) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}

func SendCodeToEmailService(cfg config.Config, code string, email string) error {
	sender := NewGmailSender("Your verification code", cfg.EmailForEmailToSend, cfg.PasswordForEmailToSend)
	subject := "SmartLab"
	content := fmt.Sprintf("Hello, it's SmartLab\n And that is your authorization code : %s", code)
	to := []string{email}
	err := sender.SendEmail(subject, content, to, nil, nil)
	return err
}
