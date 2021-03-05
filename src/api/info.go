package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/a-cordier/httpt/util"
)

type Info struct {
	HostName string `json:"hostName"`
}

var info = buildInfo()

func buildInfo() *Info {
	hostName, _ := os.Hostname()
	return &Info{
		hostName,
	}
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(util.WithJson(w)).Encode(info)
}

func RegisterInfoHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /info", "Get infos")
	server.HandleFunc("/info", util.LogTime(infoHandler))
}
