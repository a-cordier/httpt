package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/a-cordier/httpt/util"
)

type requestData struct {
	Method  string      `json:"method"`
	URI     string      `json:"uri"`
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
}

func readBody(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func reflectHandler(w http.ResponseWriter, r *http.Request) {
	body, err := readBody(r)
	if err != nil {
		newStatus(500).write(w)
	}
	json.NewEncoder(util.WithJson(w)).Encode(&requestData{
		r.Method,
		r.RequestURI,
		r.Header,
		body,
	})
}

func RegisterReflectHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /reflect", "Prints request as JSON")
	server.HandleFunc("/api/reflect", util.LogTime(reflectHandler))
}
