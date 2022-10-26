package main

import (
	"lesson15/config"
	"lesson15/internal/server"
	"lesson15/internal/services"
	"log"
	"net"
	"net/http"
)

type asd struct {
}

func (a asd) ServeHTTP(r http.ResponseWriter, req *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	asd := asd{}
	service := services.NewService()

	newServer := server.NewServer(mux, service)
	newServer.Init()

	getConfig, err := config.GetConfig()
	if err != nil {
		log.Println("Ne poluchilos poluchit nastroyki")
		return
	}

	address := net.JoinHostPort(getConfig.Host, getConfig.Port)
	log.Println(address)
	err = http.ListenAndServe(address, asd)
	if err != nil {
		log.Println(err)
	}

}
