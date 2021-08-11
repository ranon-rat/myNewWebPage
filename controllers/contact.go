package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ranon-rat/portofolio/types"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	r.ParseForm()
	email := r.Form.Get("email")
	subject := r.Form.Get("subject")
	message := r.Form.Get("body")
	webHookRequest(email, subject, message, r.RemoteAddr)
}

func webHookRequest(email, subject, message, ip string) {
	params, _ := json.Marshal(types.Message{
		Content: "new contact",
		Embeds: []types.Embed{
			{
				Title: "Contact from`" + email + "`",
				Fields: []types.Field{
					{
						Name:  "email",
						Value: "```" + email + "```",
					},
					{
						Name:  "subject",
						Value: "```" + subject + "```",
					},
					{
						Name:  "message",
						Value: "```" + message[:len(message)%500] + "```",
					},
					{
						Name:  "date",
						Value: "```" + time.Now().Format("2006-01-02 15:04:05") + "```",
					},
					{
						Name:  "ip",
						Value: "```" + ip + "```",
					},
				}},
		}})

	responseBody := bytes.NewBuffer(params)

	_, err := http.Post(types.Webhook, "application/json", responseBody)
	if err != nil {
		log.Println(err)
	}

}
