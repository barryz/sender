package model

import (
	"fmt"
)

type Sms struct {
	Tos     string `json:"tos"`
	Content string `json:"content"`
}

type Mail struct {
	Tos     string `json:"tos"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type Slack struct {
	Status  string `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (this *Sms) String() string {
	return fmt.Sprintf(
		"<Tos:%s, Content:%s>",
		this.Tos,
		this.Content,
	)
}

func (this *Mail) String() string {
	return fmt.Sprintf(
		"<Tos:%s, Subject:%s, Content:%s>",
		this.Tos,
		this.Subject,
		this.Content,
	)
}

func (this *Slack) String() string {
	return fmt.Sprintf(
		"<Title:%s, Status:%s, Content:%s>",
		this.Title,
		this.Status,
		this.Content,
	)
}
