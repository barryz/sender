package cron

import (
	"github.com/barryz/sender/g"
	"github.com/barryz/sender/model"
	"github.com/barryz/sender/proc"
	"github.com/barryz/sender/redis"
	"github.com/toolkits/net/httplib"
	"log"
	"time"
)

func ConsumeMail() {
	queue := g.Config().Queue.Mail
	for {
		L := redis.PopAllMail(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		MailListChan <- L
		SendMailList(L)
	}
}

func SendMailList(L []*model.Mail) {
	for _, mail := range L {
		MailWorkerChan <- 1
		go SendMail(mail)
	}
}

func SendMail(mail *model.Mail) {
	defer func() {
		<-MailWorkerChan
	}()

	url := g.Config().Api.Mail
	r := httplib.Post(url).SetTimeout(5 * time.Second, 2 * time.Minute)
	r.Param("tos", mail.Tos)
	r.Param("subject", mail.Subject)
	r.Param("content", mail.Content)
	resp, err := r.String()
	if err != nil {
		log.Println(err)
	}

	proc.IncreMailCount()

	if g.Config().Debug {
		log.Println("==mail==>>>>", mail)
		log.Println("<<<<==mail==", resp)
	}

}
