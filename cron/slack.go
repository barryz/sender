package cron

import (
	"encoding/json"
	"log"
	"time"

	"sender/g"
	"sender/model"
	"sender/proc"
	"sender/redis"

	"github.com/toolkits/net/httplib"
)

func ConsumeSlack() {
	queue := g.Config().Queue.Slack
	for {
		L := redis.PopAllSlack(queue)
		if len(L) == 0 {
			time.Sleep(200 * time.Millisecond)
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
	slackContent, err := json.Marshal(slack.Content)
	if err != nil {
		log.Printf("[ERROR]: parse slack alarm content fail due to %s", err.Error())
		return
	}
	r := httplib.Post(url).SetTimeout(5*time.Second, 2*time.Minute)
	r.Param("channel", slack.Channel)
	r.Param("content", string(slackContent))
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
