package proc

import (
	"sync/atomic"
)

var smsCount, mailCount, slackCount uint32

func GetSmsCount() uint32 {
	return atomic.LoadUint32(&smsCount)
}

func GetMailCount() uint32 {
	return atomic.LoadUint32(&mailCount)
}

func GetSlackCount() uint32 {
	return atomic.LoadUint32(&slackCount)
}

func IncreSmsCount() {
	atomic.AddUint32(&smsCount, 1)
}

func IncreMailCount() {
	atomic.AddUint32(&mailCount, 1)
}

func IncreSlackCount() {
	atomic.AddUint32(&slackCount, 1)
}
