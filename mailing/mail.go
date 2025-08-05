package main

import "net/smtp"

func SendMail(mail string, password string, msg string) {
	auth := smtp.PlainAuth(
		"",
		mail,
		password,
		"smpt.gmail.com",
	)
	smtp.SendMail(
		"smpt.gmail.com:587",
		auth,
		mail,
		[]string{"suryanshverma.dev.official@gmail.com"},
		[]byte(msg),
	)
}
