package main

import (
	"lesson15/config"
	"lesson15/internal/server"
	"lesson15/internal/services"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	service := services.NewService()

	newServer := server.NewServer(mux, service)
	
	getConfig, err := config.GetConfig()
	if err != nil {
		log.Println("Ne poluchilos poluchit nastroyki")
		return
	}

	address := net.JoinHostPort(getConfig.Host, getConfig.Port)
	log.Println(address)
	err = http.ListenAndServe(address, newServer)
	if err != nil {
		log.Println(err)
	}

}
