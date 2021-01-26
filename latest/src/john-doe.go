package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type JohnDoe struct {
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Padding   string `json:"padding"`
}

const (
	oneHundredBytesJohnDoe = `{
		"name": "John Doe",
		"firstName": "John",
		"lastName": "Doe",
		"email": "john@doe.dev",
		"padding": "******"
	  }`
)

func parseJohnDoe() JohnDoe {
	doe := &JohnDoe{}
	json.Unmarshal([]byte(oneHundredBytesJohnDoe), doe)
	return *doe
}

func createJohnDoes(times int) []JohnDoe {
	does := make([]JohnDoe, times)
	doe := parseJohnDoe()
	for i := 0; i < times; i++ {
		doe := doe
		if times > 1 && i < times-1 {
			doe.Padding += "**" // ensure John weights its 100 bytes
		}
		does[i] = doe
	}
	return does
}

func johnDoeHandler(w http.ResponseWriter, r *http.Request) {
	times := parseIntQuery(r.URL, "times", 1)
	does := createJohnDoes(times)
	json.NewEncoder(withJson(w)).Encode(does)
}

func registerJohnDoesHandler(server *http.ServeMux) {
	log.Printf("%s (%s)", "Registering [*] /john-doe?times={count}", "Returns 100 bytes of John Doe times {count}")
	server.HandleFunc("/john-doe", httpLogger(johnDoeHandler))
}
