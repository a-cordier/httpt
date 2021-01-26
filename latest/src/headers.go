package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func headeReflectHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(withJson(w)).Encode(r.Header)
}

func registerHeadersHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /headers", "Prints request headers as JSON")
	server.HandleFunc("/headers", httpLogger(headeReflectHandler))
}
