package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const EMAIL_HOST = "localhost"
const EMAIL_PORT = 1025
const EMAIL_AUTH_USERNAME = ""
const EMAIL_AUTH_PASSWORD = ""
const EMAIL_FROM = "notifications@origins.link"

// Sends an email
func sendEmail(to []string, subject string, msg []byte) error {
	var auth smtp.Auth

	e := &email.Email{
		To:      to,
		From:    EMAIL_FROM,
		Subject: subject,
		Text:    msg,
	}

	if EMAIL_AUTH_USERNAME != "" {
		auth = smtp.PlainAuth("", EMAIL_AUTH_USERNAME, EMAIL_AUTH_PASSWORD, EMAIL_HOST)
	}

	addr := fmt.Sprintf("%s:%d", EMAIL_HOST, EMAIL_PORT)

	return e.Send(addr, auth)
}
