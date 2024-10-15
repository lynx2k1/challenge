package mail

import (
    "net/smtp"
)

func SendEmail(to string, subject string, body string) error {
    from := "your-email@example.com"
    password := "your-email-password"

    // Configuração do servidor SMTP
    smtpHost := "smtp.example.com"
    smtpPort := "587"

    // Monta a mensagem
    message := []byte("Subject: " + subject + "\r\n" +
        "\r\n" + body)

    // Autenticação
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Envia o e-mail
    return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
}
