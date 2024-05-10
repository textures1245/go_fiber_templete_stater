package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (e *Email) Send() error {
	serviceProvider := fmt.Sprintf("%s:%d", viper.GetString("MAILERSEND_SERVICE"), viper.GetInt("MAILERSEND_PORT"))

	auth := smtp.PlainAuth(
		"",
		viper.GetString("MAILERSEND_USERNAME"),
		viper.GetString("MAILERSEND_PASSWORD"),
		viper.GetString("MAILERSEND_SERVICE"),
	)

	to := []string{e.To}
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		e.From, e.To, e.Subject, e.Body))

	conn, err := smtp.Dial(serviceProvider)
	if err != nil {
		return err
	}

	// Upgrade to a secure connection
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         viper.GetString("MAILERSEND_SERVICE"),
	}
	conn.StartTLS(tlsConfig)

	if err := smtp.SendMail(serviceProvider, auth, e.From, to, msg); err != nil {
		return err
	}

	return nil
}
