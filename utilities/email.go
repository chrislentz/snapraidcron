package utilities

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func SendSmtp(host string, port string, username string, password string, to string, from string, subject string, body string) error {
	toList := strings.Split(to, ",")

	msg := fmt.Sprintf("To: %s\nFrom: %s\nSubject: %s\n\n%s", to, from, subject, body)

	smtpBody := []byte(msg)

	auth := smtp.PlainAuth("", username, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, smtpBody)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
