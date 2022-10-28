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
	err := execute()
	if err != nil {
		log.Println(err)
		return
	}
}

func execute() error {
	mux := http.NewServeMux()

	service := services.NewService()

	newServer := server.NewServer(mux, service)
	newServer.Init()

	getConfig, err := config.GetConfig()
	if err != nil {
		log.Println("Ne poluchilos poluchit nastroyki")
		return err
	}

	address := net.JoinHostPort(getConfig.Host, getConfig.Port)

	srv := http.Server{
		Addr:    address,
		Handler: newServer,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	return nil

}
