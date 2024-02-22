package main

import "github.com/gabrielsscti/event-driven-web-server/pkg/webserver"

func main() {
	server := webserver.CreateSingleThreadWebServer(webserver.TCP, 8080)
	server.ListenAndServe()
}
