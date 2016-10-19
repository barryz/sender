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

func CounsumeSlack() {
	queue := g.Config().Queue.Mail
	for {
		L := redis.PopAllMail(queue)
		if len(L) == 0 {
			time.Sleep(time.Microsecond * 200)
			continue
		}
		SendSlackList(L)
	}
}

func SendSlackList(L []*model.Mail) {
	for _, slack := range L {
		SlackWorkerChan <- 1
		go SendSlack(slack)
	}
}

func SendSlack(slack *model.Mail) {
	defer func() {
		<-SlackWorkerChan
	}()

	url := g.Config().Api.Slack
	r := httplib.Post(url).SetTimeout(5 * time.Second, 2 * time.Minute)
	r.Param("tos", slack.Tos)
	r.Param("subject", slack.Subject)
	r.Param("content", slack.Content)
	resp, err := r.String()
	if err != nil {
		log.Println(err)
	}

	proc.IncreSlackCount()

	if g.Config().Debug {
		log.Println("==slack==>>>>", slack)
		log.Println("<<<<===slack==", resp)
	}

}