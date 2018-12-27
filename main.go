package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var channelString = make(chan string)

const Host string = "0.0.0.0"

func CreateProxyServer(fromPort, toPort string) {
	redirectToURL, _ := url.Parse(fmt.Sprintf("http://%s:%s", Host, toPort))
	reverseProxy := httputil.NewSingleHostReverseProxy(redirectToURL)
	http.Handle("/", reverseProxy)
	channelString <- fmt.Sprintf("start forward port: %s -> %s", fromPort, toPort)
	error := http.ListenAndServe(fmt.Sprintf("%s:%s", Host, fromPort), nil)
	channelString <- error.Error()
}

func main() {
	go CreateProxyServer("8080", "4480")

	for {
		fmt.Println(<-channelString)
	}
}
