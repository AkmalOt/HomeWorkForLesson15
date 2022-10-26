package handlers

import (
	"encoding/json"
	"io"
	"lesson15/internal/models"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Calculation(response http.ResponseWriter, request *http.Request) {

	quaries := request.URL.Query()
	FirstNum := quaries.Get("first_num")
	SecondNum := quaries.Get("second_num")
	Operation := quaries.Get("operation")

	FirstNumInt, err := strconv.Atoi(FirstNum)
	if err != nil {
		log.Println(err)
		return
	}
	SecondNumInt, err := strconv.Atoi(SecondNum)
	if err != nil {
		log.Println(err)
		return
	}
	var ActionOfCulc int
	switch {
	case Operation == "+":
		ActionOfCulc = FirstNumInt + SecondNumInt
	case Operation == "-":
		ActionOfCulc = FirstNumInt - SecondNumInt
	case Operation == "*":
		ActionOfCulc = FirstNumInt * SecondNumInt
	case Operation == "/":
		ActionOfCulc = FirstNumInt / SecondNumInt
	}
	config := models.CalcResult{
		FirstNum:  FirstNumInt,
		SecondNum: SecondNumInt,
		Operation: Operation,
		Result:    ActionOfCulc,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	//даем пользователю на экран инфу
	_, err = response.Write(data)
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.OpenFile("./Result.json", os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	contentJson, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	var History []models.CalcResult
	err = json.Unmarshal(contentJson, &History)
	if err != nil {
		log.Println(err)
	}

	History = append(History, config)

	bytes, err := json.MarshalIndent(History, "", "  ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("Result.json", bytes, 0777)
	if err != nil {
		log.Println(err)
	}
}

//const InfoFile = "Result.json"

func GetHistory(response http.ResponseWriter, request *http.Request) {

	var History []models.CalcResult
	file, err := os.OpenFile("./Result.json", os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	contentJson, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(contentJson, &History)
	if err != nil {
		log.Println(err)
	}

	_, err = response.Write(contentJson)
	if err != nil {
		log.Println("Не получилось")
		return
	}
}

func CleanHistory(response http.ResponseWriter, request *http.Request) {
	file, err := os.OpenFile("./Result.json", os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Println(err)
		return
	}
	err = file.Truncate(0)
	file.Close()
}
