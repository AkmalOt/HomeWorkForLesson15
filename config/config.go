package config

import (
	"encoding/json"
	"io"
	"lesson15/internal/models"
	"log"
	"os"
)

func GetConfig() (*models.Config, error) {
	//чтение и дессериализация конфигов
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var config models.Config

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &config, err
}
