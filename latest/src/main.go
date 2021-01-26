package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const banner = `

@@@  @@@ @@@@@@@ @@@@@@@ @@@@@@@  @@@@@@@
@@!  @@@   @@!     @@!   @@!  @@@   @@!  
@!@!@!@!   @!!     @!!   @!@@!@!    @!!  
!!:  !!!   !!:     !!:   !!:        !!:  
 :   : :    :       :     :          :  

https://github.com/adeo/dockerfiles-collection/http-test-server  

`

func getEnvOrDefault(varName string, defaultVal string) string {
	val := os.Getenv(varName)
	if val == "" {
		return defaultVal
	}
	return val
}

func parseIntQuery(URL *url.URL, name string, or int) int {
	param := URL.Query().Get(name)
	if param == "" {
		return or
	}
	val, err := strconv.Atoi(param)
	if err != nil {
		return or
	}
	return val
}

func withJson(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
}

func httpLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		log.Printf("[%s] %s took %s", r.Method, r.URL.String(), elapsed)
	}
}

func main() {
	log.Print(banner)
	server := http.NewServeMux()
	port := getEnvOrDefault("PORT", "8080")

	registerStatusHandlers(server)
	registerHeadersHandler(server)
	registerJohnDoesHandler(server)
	registerWaitHandler(server)

	log.Printf("%s %s\n\n", "Starting httpt on port", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}
