package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"sync"

	"gopkg.in/gomail.v2"
)

//go:embed templates
var mailTemplates embed.FS

type MailHelper struct {
	Hostname string
	Port     int
	Username string
	Password string
	Dialer   *gomail.Dialer
}

func NewMailer(hostname, username, password string, port int) *MailHelper {
	return &MailHelper{
		Hostname: hostname,
		Port:     port,
		Username: username,
		Password: password,
	}
}

type MailMessage struct {
	To        string
	From      string
	Subject   string
	EmailAddr string
	Message   string
}

func (m *MailHelper) NewDialer() {
	m.Dialer = gomail.NewDialer(m.Hostname, m.Port, m.Username, m.Password)
}

func (m *MailHelper) NewMessage(toAddr, fromAddr, subject, emailAddr, msgContent string) *MailMessage {
	var msg = &MailMessage{
		To:        toAddr,
		From:      fromAddr,
		Subject:   subject,
		EmailAddr: emailAddr,
		Message:   msgContent,
	}

	return msg
}

func (m *MailHelper) SendMessage(msg *MailMessage, tmpl string) error {

	type MailData struct {
		Subject string
		Email   string
		Message string
	}

	var md MailData

	md.Subject = msg.Subject
	md.Email = msg.EmailAddr
	md.Message = msg.Message

	t, err := template.New("email").ParseFS(mailTemplates, "templates/"+tmpl)

	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = t.ExecuteTemplate(subject, "subject", md)

	if err != nil {
		return err
	}

	textBody := new(bytes.Buffer)

	err = t.ExecuteTemplate(textBody, "textBody", md)

	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)

	err = t.ExecuteTemplate(htmlBody, "htmlBody", md)

	if err != nil {
		return err
	}

	msgToSend := gomail.NewMessage()

	msgToSend.SetHeader("To", msg.To)
	msgToSend.SetHeader("From", msg.From)
	msgToSend.SetHeader("Subject", subject.String())
	msgToSend.SetBody("text/plain", textBody.String())
	msgToSend.AddAlternative("text/html", htmlBody.String())
	err = m.Dialer.DialAndSend(msgToSend)

	if err != nil {
		return err
	}

	return nil
}

func (m *MailHelper) ListenForMail(msgChan <-chan *MailMessage, errChan chan<- error, doneChan <-chan bool, wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		select {
		case msg := <-msgChan:
			err := m.SendMessage(msg, "contact-form.gohtml")
			if err != nil {
				fmt.Println(err)
				errChan <- err
			}
		case <-doneChan:
			close(errChan)
			return

		}
	}

}
