package redis

import (
	"encoding/json"
	"log"
	"sender/model"

	"github.com/garyburd/redigo/redis"
)

func PopAllSlack(queue string) []*model.Slack {
	ret := []*model.Slack{}

	rc := ConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var slack model.Slack
		err = json.Unmarshal([]byte(reply), &slack)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &slack)
	}

	return ret
}

func PopAllSms(queue string) []*model.Sms {
	ret := []*model.Sms{}

	rc := ConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var sms model.Sms
		err = json.Unmarshal([]byte(reply), &sms)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &sms)
	}

	return ret
}

func PopAllMail(queue string) []*model.Mail {
	ret := []*model.Mail{}

	rc := ConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var mail model.Mail
		err = json.Unmarshal([]byte(reply), &mail)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &mail)
	}

	return ret
}
