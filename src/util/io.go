package util

import (
	"net/http"
	"net/url"
	"strconv"
)

func ParseIntQuery(URL *url.URL, name string, or int) int {
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

func WithJson(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
}
