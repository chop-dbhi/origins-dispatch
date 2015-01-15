package main

import (
	"net"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

// Sends an email
func sendEmail(to []string, subject string, msg []byte) error {
	var auth smtp.Auth

	from := viper.GetString("from")

	e := &email.Email{
		To:      to,
		From:    from,
		Subject: subject,
		Text:    msg,
	}

	addr := viper.GetString("smtp_addr")
	user := viper.GetString("smtp_user")

	if user != "" {
		password := viper.GetString("smtp_password")

		host, _, err := net.SplitHostPort(addr)

		if err != nil {
			return err
		}

		auth = smtp.PlainAuth("", user, password, host)
	}

	return e.Send(addr, auth)
}
