package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func isURL(str string) (*url.URL, bool) {
	url, err := url.Parse(str)
	return url, err == nil && url.Scheme != "" && url.Host != ""
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
