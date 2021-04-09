package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-cordier/httpt/util"
)

func headerReflectHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(util.WithJson(w)).Encode(r.Header)
}

func RegisterHeadersHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /headers", "Prints request headers as JSON")
	server.HandleFunc("/headers", util.LogTime(headerReflectHandler))
}
