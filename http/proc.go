package http

import (
	"fmt"
	"net/http"
	"sender/proc"
)

func configProcRoutes() {

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("sms:%v, mail:%v", proc.GetSmsCount(), proc.GetMailCount())))
	})

}
