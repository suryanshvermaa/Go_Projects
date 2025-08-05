package main

import (
	"log"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func SendMail(mail string, password string, msg string) {
	auth := smtp.PlainAuth(
		"",
		mail,
		password,
		"smtp.gmail.com",
	)
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		mail,
		[]string{"suryanshverma.dev.official@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func SendMailByGoMail(mail string, password string, msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", mail)
	m.SetHeader("To", "suryanshverma.dev.official@gmail.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("./hello.txt")

	d := gomail.NewDialer("smtp.gmail.com", 587, mail, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
