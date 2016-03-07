package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	proxyURL, _ := url.Parse("https://byt.tl")

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	http.ListenAndServe(":8080", proxy)
}
