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

func infoHandler(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	json.NewEncoder(util.WithJson(w)).Encode(&Info{name})
}

func RegisterInfoHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /info", "Get infos")
	server.HandleFunc("/info", util.HttpLogger(infoHandler))
}
