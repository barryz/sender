package cron

import (
	"github.com/barryz/sender/g"
	"github.com/barryz/sender/model"
)

var (
	SmsWorkerChan  chan int
	MailWorkerChan chan int
	SlackWorkerChan chan int
	MailListChan chan []*model.Mail
)

func InitWorker() {
	workerConfig := g.Config().Worker
	SmsWorkerChan = make(chan int, workerConfig.Sms)
	MailWorkerChan = make(chan int, workerConfig.Mail)
	SlackWorkerChan = make(chan int, workerConfig.Slack)
	MailListChan = make(chan []*model.Mail, 1)
}
