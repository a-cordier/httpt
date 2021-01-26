package main

import (
	"log"
	"net/http"
	"time"
)

func waitHandler(w http.ResponseWriter, r *http.Request) {
	var seconds = parseIntQuery(r.URL, "seconds", 0)
	time.Sleep(time.Duration(seconds) * time.Second)
	newStatus(http.StatusOK).write(w)
}

func registerWaitHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /wait?seconds={count}", "Waits during {count} seconds, then returns 200 OK")
	server.HandleFunc("/wait", httpLogger(waitHandler))
}
