package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-cordier/httpt/util"
)

func headeReflectHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(util.WithJson(w)).Encode(r.Header)
}

func RegisterHeadersHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /headers", "Prints request headers as JSON")
	server.HandleFunc("/headers", util.HttpLogger(headeReflectHandler))
}
