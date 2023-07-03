package utils

import (
	"api-public-platform/config"
	"fmt"

	"gopkg.in/gomail.v2"
)

type Mailer interface {
	SendMail(to, subject, body string) error
}

type QQMailer struct {
	Host     string
	Port     int
	Username string
	Password string
}

type GmailMailer struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (m *QQMailer) SendMail(to, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("send email failed: %v", err)
	}
	return nil
}

func (m *GmailMailer) SendMail(to, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("send email failed: %v", err)
	}
	return nil
}

type MailSender struct {
	mailer Mailer
}

func NewMailSender(mailer Mailer) *MailSender {
	return &MailSender{mailer: mailer}
}

func NewQQMailSender() *MailSender {
	qqMailer := QQMailer{
		Username: config.ServerCfg.Email.QQ.Username,
		Password: config.ServerCfg.Email.QQ.Password,
		Host:     config.ServerCfg.Email.QQ.Host,
		Port:     config.ServerCfg.Email.QQ.Port,
	}
	return NewMailSender(&qqMailer)
}

func NewQQMailSenderCustom(host string, port int, username, password string) *MailSender {
	qqMailer := QQMailer{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
	return NewMailSender(&qqMailer)
}

func NewGmailSender() *MailSender {
	gmailMailer := GmailMailer{
		Username: config.ServerCfg.Email.Gmail.Username,
		Password: config.ServerCfg.Email.Gmail.Password,
		Host:     config.ServerCfg.Email.Gmail.Host,
		Port:     config.ServerCfg.Email.Gmail.Port,
	}
	return NewMailSender(&gmailMailer)
}

func (s *MailSender) SendMail(to, subject, body string) error {
	return s.mailer.SendMail(to, subject, body)
}
