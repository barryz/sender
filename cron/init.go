package cron

import (
	"sender/g"
)

var (
	SmsWorkerChan   chan int
	MailWorkerChan  chan int
	SlackWorkerChan chan int
)

func InitWorker() {
	workerConfig := g.Config().Worker
	SmsWorkerChan = make(chan int, workerConfig.Sms)
	MailWorkerChan = make(chan int, workerConfig.Mail)
	SlackWorkerChan = make(chan int, workerConfig.Slack)
}
