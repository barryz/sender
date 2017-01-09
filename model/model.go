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
	Channel string        `json:"channel"`
	Content *SlackContent `json:"content"`
}

type SlackContent struct {
	EndPoint     string `json:"endpoint"`
	Note         string `json:"note"`
	Status       string `json:"status"`
	Priority     string `json:"priority"`
	Metric       string `json:"metric"`
	CurrentValue string `json:"current_value"`
	Expression   string `json:"expr"`
	AlarmCount   string `json:"alarm_count"`
	TriggerTime  string `json:"trigger_time"`
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
		"<Channel: %s, Content:%v>",
		this.Channel,
		this.Content,
	)
}
