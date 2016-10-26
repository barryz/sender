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

func ConsumeSlack() {
	queue := g.Config().Queue.Slack
	for {
		L := redis.PopAllSlack(queue)
		if len(L) == 0 {
			time.Sleep(200 * time.Microsecond)
			continue
		}
		SendSlackList(L)
	}
}

func SendSlackList(L []*model.Slack) {
	for _, slack := range L {
		SlackWorkerChan <- 1
		go SendSlack(slack)
	}
}

func SendSlack(slack *model.Slack) {
	defer func() {
		<-SlackWorkerChan
	}()

	url := g.Config().Api.Slack
	r := httplib.Post(url).SetTimeout(5 * time.Second, 2 * time.Minute)
	r.Param("title", slack.Title)
	r.Param("status", slack.Status)
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