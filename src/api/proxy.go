package api

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/a-cordier/httpt/util"
)

func decodeUrl(r *http.Request) (*url.URL, error) {
	encoded := r.URL.Query().Get("target")
	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		return nil, err
	}
	return url.Parse(decoded)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	url, err := decodeUrl(r)
	if err != nil {
		newStatus(400).write(w)
	}
	r.URL = url
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host
	u, _ := url.Parse("/")
	httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
}

func RegisterProxyHandler(server *http.ServeMux) {
	log.Printf("%s %s (%s)", "Registering [*]", "/proxy?target={target}", "Proxy request to encoded url target")
	server.HandleFunc("/proxy", util.LogTime(proxyHandler))
}
