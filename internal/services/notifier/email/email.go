package email

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"io"
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
	"strings"
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

	// Создаем HTML сообщение с встроенными изображениями
	htmlContent := message

	// Если есть изображения, встраиваем их в HTML
	for i, file := range files {
		contentID := fmt.Sprintf("image%d", i+1)

		// Добавляем встроенное изображение
		attachment, err := AttachInlineImageFromURL(msg, file, contentID)
		if err != nil {
			fmt.Printf("couldn't attach inline image %s: %v\n", file, err)
			continue
		}

		msg.Attachments = append(msg.Attachments, attachment)

		// Если изображение не встроено в HTML, добавляем его в конец сообщения
		if !strings.Contains(htmlContent, "cid:"+contentID) {
			htmlContent += fmt.Sprintf("<br><img src=\"cid:%s\" alt=\"Image %d\">", contentID, i+1)
		}
	}

	if m.usePlainText {
		msg.Text = []byte(message)
	} else {
		msg.HTML = []byte(htmlContent)
	}

	return msg
}

func (m *Mail) Send(ctx context.Context, subject, message string, receivers []string, images ...string) error {
	htmlMessage := message

	if !m.usePlainText && len(images) > 0 {
		// Создаем массив для отслеживания использованных изображений
		usedImages := make([]bool, len(images))

		// Обрабатываем индивидуальные плейсхолдеры для изображений
		for i := range images {
			imagePlaceholder := fmt.Sprintf("{{IMAGE:%d}}", i)
			if strings.Contains(htmlMessage, imagePlaceholder) {
				imageTag := fmt.Sprintf("<img src=\"cid:image%d\" alt=\"Image %d\" style=\"width:300px;height:auto;margin:10px 20px;\">", i+1, i+1)
				htmlMessage = strings.Replace(htmlMessage, imagePlaceholder, imageTag, -1)
				usedImages[i] = true
			}
		}

		// Обрабатываем общий плейсхолдер {{IMAGES}}
		if strings.Contains(htmlMessage, "{{IMAGES}}") {
			var remainingImagesHTML string
			for i, used := range usedImages {
				if !used {
					remainingImagesHTML += fmt.Sprintf("<img src=\"cid:image%d\" alt=\"Image %d\" style=\"width:300px;height:auto;margin:10px 20px;\"><br>", i+1, i+1)
				}
			}
			htmlMessage = strings.Replace(htmlMessage, "{{IMAGES}}", remainingImagesHTML, 1)
		}
	}

	// Создаем email напрямую
	msg := &email.Email{
		To:      validEmails(receivers),
		From:    m.senderAddress,
		Subject: subject,
		Headers: textproto.MIMEHeader{},
	}

	// Прикрепляем изображения
	for i, imageURL := range images {
		contentID := fmt.Sprintf("image%d", i+1)
		attachment, err := AttachInlineImageFromURL(msg, imageURL, contentID)
		if err != nil {
			fmt.Printf("couldn't attach inline image %s: %v\n", imageURL, err)
			continue
		}
		msg.Attachments = append(msg.Attachments, attachment)
	}

	// Устанавливаем содержимое сообщения
	if m.usePlainText {
		msg.Text = []byte(message)
	} else {
		msg.HTML = []byte(htmlMessage)
	}

	// Отправляем email
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

// AttachInlineImageFromURL загружает изображение по URL и добавляет его как встроенное изображение
func AttachInlineImageFromURL(msg *email.Email, fileURL string, contentID string) (*email.Attachment, error) {
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

	// Проверяем, что это изображение
	if !strings.HasPrefix(contentType, "image/") {
		// Если это не изображение, прикрепляем как обычный файл
		return msg.Attach(resp.Body, name, contentType)
	}

	// Читаем содержимое файла
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Создаем встроенное изображение с Content-ID
	attachment := &email.Attachment{
		Filename:    name,
		ContentType: contentType,
		Header:      textproto.MIMEHeader{},
		Content:     data,
	}

	attachment.Header.Set("Content-ID", "<"+contentID+">")
	attachment.Header.Set("Content-Disposition", "inline")

	return attachment, nil
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
