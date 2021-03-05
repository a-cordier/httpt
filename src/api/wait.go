package api

import (
	"log"
	"net/http"
	"time"

	"github.com/a-cordier/httpt/util"
)

func waitHandler(w http.ResponseWriter, r *http.Request) {
	var seconds = util.ParseIntQuery(r.URL, "seconds", 0)
	time.Sleep(time.Duration(seconds) * time.Second)
	newStatus(http.StatusOK).write(w)
}

func RegisterWaitHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /wait?seconds={count}", "Waits during {count} seconds, then returns 200 OK")
	server.HandleFunc("/wait", util.LogTime(waitHandler))
}
