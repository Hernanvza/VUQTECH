package utils

import (
	"net/smtp"
)

// SendEmail envía un correo electrónico utilizando un servidor SMTP
func SendEmail(to string, subject string, body string) error {
	from := "hernan.valenzuela@pudupaw.cl"
	password := "Eengeo9ej5wo@"
	smtpHost := "c2670935.ferozo.com"
	smtpPort := "995"

	// Construir el mensaje
	message := []byte("Subject: " + subject + "\n\n" + body)

	// Autenticación SMTP
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Enviar el correo
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
}
