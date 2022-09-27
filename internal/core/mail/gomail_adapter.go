package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	customerror "github.com/eneskzlcn/softdare/internal/error"
	gomail "gopkg.in/mail.v2"
)

type GomailServiceAdapter struct {
	config config.MailService
	logger logger.Logger
}

func NewGomailServiceAdapter(config config.MailService, logger logger.Logger) *GomailServiceAdapter {
	if logger == nil {
		fmt.Println(customerror.NilLogger)
		return nil
	}
	return &GomailServiceAdapter{config: config, logger: logger}
}
func (g *GomailServiceAdapter) SendTextMail(to, subject, content string) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", g.config.SenderMail)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", content)

	// Settings for SMTP server

	d := gomail.NewDialer(g.config.Host, g.config.Port, g.config.SenderMail, g.config.SenderPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		g.logger.Error("error when sending the email.")
		return err
	}
	return nil
}
