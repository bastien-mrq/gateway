package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/charmbracelet/log"
)

func logRequest(r *http.Request) {
	log.Infof("Received %s request for %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

func proxyRequest(target string) http.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL %s: %v", target, err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		proxy.ServeHTTP(w, r)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		http.Error(w, "Not found", http.StatusNotFound)
	})

	log.Print("Gateway running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
