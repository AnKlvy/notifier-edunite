package email

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"mime"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"os"
	"path"
	"path/filepath"
)

// Mail struct holds necessary data to send emails.
type Mail struct {
	usePlainText      bool
	senderAddress     string
	smtpHostAddr      string
	smtpAuth          smtp.Auth
	receiverAddresses []string
}

// New returns a new instance of a Mail notification service.
func InitEmail() *Mail {
	senderAddress, password, smtpHostAddress := os.Getenv("SENDER_EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("SMTP_HOST_ADDRESS")
	mail := &Mail{
		usePlainText:      false,
		senderAddress:     senderAddress,
		smtpHostAddr:      smtpHostAddress,
		receiverAddresses: []string{},
	}

	host, _, err := net.SplitHostPort(smtpHostAddress)
	if err != nil {
		_ = fmt.Errorf("error spliting host/port: %v", err)
	}
	log.Println("host:", host)

	mail.smtpAuth = smtp.PlainAuth("", senderAddress, password, host)
	return mail
}

func (m *Mail) newEmail(subject, message string, receivers []string, files ...string) *email.Email {
	msg := &email.Email{
		To:      validEmails(receivers),
		From:    m.senderAddress,
		Subject: subject,
		Headers: textproto.MIMEHeader{},
	}

	for _, file := range files {
		attachment, err := AttachFromURL(msg, file)
		if err != nil {
			fmt.Printf("couldn't attach file %s: %v\n", file, err)
			continue
		}
		msg.Attachments = append(msg.Attachments, attachment)
	}
	if m.usePlainText {
		msg.Text = []byte(message)
	} else {
		msg.HTML = []byte(message)
	}
	return msg
}

func (m *Mail) Send(ctx context.Context, subject, message string, receivers []string, images ...string) error {

	msg := m.newEmail(subject, message, receivers, images...)

	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		if err = msg.Send(m.smtpHostAddr, m.smtpAuth); err != nil {
			err = fmt.Errorf("send email: %w", err)
		}
	}

	return err
}

// TODO ограничить вес файла (валидация)
func AttachFromURL(msg *email.Email, fileURL string) (*email.Attachment, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status downloading file: %v", resp.Status)
	}

	// Выдёргиваем имя файла из URL
	_, name := path.Split(fileURL)

	// Определяем content-type
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = mime.TypeByExtension(filepath.Ext(name))
	}

	return msg.Attach(resp.Body, name, contentType)
}

func validEmails(to []string) []string {
	var validEmails []string

	for _, addrStr := range to {
		addr, err := mail.ParseAddress(addrStr)
		if err == nil {
			validEmails = append(validEmails, addr.Address)
		} else {
			log.Printf("validEmails: error parsing email: %v\n", err)
		}
	}
	return validEmails
}
