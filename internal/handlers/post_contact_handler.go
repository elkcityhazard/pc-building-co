package handlers

import (
	"bufio"
	"errors"
	"html"
	"log"
	"net/http"
	"net/mail"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/models"
	"github.com/elkcityhazard/pc-building-company/pkg/mailer"
)

var spamList []string

func (hr *HandlerRepo) PostContactHandler(w http.ResponseWriter, r *http.Request) {

	spamF, err := os.Open("./static/data/spam-list.txt")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer spamF.Close()

	scanner := bufio.NewScanner(bufio.NewReader(spamF))

	for scanner.Scan() {
		spamList = append(spamList, string(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type Matter struct {
		models.DefaultFrontMatter
	}

	var matter models.DefaultFrontMatter

	contentFS := content.GetContentFS()

	toRead, err := contentFS.Open(filepath.Join("content", "contact.md"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = frontmatter.Parse(bufio.NewReader(toRead), &matter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hr.Cfg.Renderer.SetStringMapEntry("PageTitle", html.EscapeString(matter.Title))
	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString(matter.Description))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", matter)
	hr.Cfg.Renderer.SetDataMapEntry("Content", hr.Cfg.MarkdownData["contact.md"])

	err = r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vals := &url.Values{}
	formErrs := make(map[string][]string)

	email := r.Form.Get("email")
	username := r.Form.Get("username")
	msg := r.Form.Get("message")

	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	username = strings.TrimSpace(username)
	username = strings.ToLower(username)

	msg = strings.TrimSpace(msg)

	vals.Add("email", html.EscapeString(email))
	vals.Add("message", html.EscapeString(msg))

	_, err = mail.ParseAddress(email)

	if err != nil {
		formErrs["email"] = append(formErrs["email"], err.Error())
		hr.Cfg.Renderer.SetDataMapEntry("FormErrors", formErrs)
		hr.Cfg.Renderer.SetDataMapEntry("Form", vals)
	}

	if strings.ToLower(username) != "" {
		hr.Cfg.Renderer.SetDataMapEntry("Form", vals)
		err = errors.New("invalid user input")
		formErrs["email"] = append(formErrs["email"], err.Error())
		err = hr.Cfg.Renderer.RenderTemplate(w, r, "contact.gohtml")
	}

	if msg == "" || strings.ToLower(msg) == "enter message..." {
		formErrs["message"] = append(formErrs["message"], "invalid message entry")
		hr.Cfg.Renderer.SetDataMapEntry("FormErrors", formErrs)
		hr.Cfg.Renderer.SetDataMapEntry("Form", vals)

	}

	//for _, v := range spamList {
	//	if strings.Contains(strings.ToLower(msg), strings.ToLower(v)) {
	//		fmt.Println(v, msg)
	//		formErrs["message"] = append(formErrs["message"], "invalid request")
	//		hr.Cfg.Renderer.SetDataMapEntry("FormErrors", formErrs)
	//		hr.Cfg.Renderer.SetDataMapEntry("Form", vals)
	//	}
	//}

	if len(formErrs["email"]) > 0 || len(formErrs["message"]) > 0 {
		err = hr.Cfg.Renderer.RenderTemplate(w, r, "contact.gohtml")
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	hr.Cfg.Renderer.SetStringMapEntry("PageTitle", html.EscapeString("Thank You"))
	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString(matter.Description))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", matter)
	hr.Cfg.Renderer.SetStringMapEntry("EmailAddress", html.EscapeString(email))
	hr.Cfg.Renderer.SetDataMapEntry("Content", hr.Cfg.MarkdownData["contact.md"])

	var mailMsg = &mailer.MailMessage{
		From:      hr.Cfg.SMTPFromAddr,
		To:        hr.Cfg.SMTPToAddr,
		Subject:   "New Submission From: " + hr.Cfg.WebsiteName,
		EmailAddr: html.EscapeString(email),
		Message:   html.EscapeString(msg),
	}

	hr.Cfg.MailMsgChan <- mailMsg

	err = hr.Cfg.Renderer.RenderTemplate(w, r, "success.gohtml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
