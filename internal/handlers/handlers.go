package handlers

import (
	"encoding/json"
	"lesson15/internal/models"
	"lesson15/internal/services"
	"log"
	"net/http"
)

func Calculation(response http.ResponseWriter, request *http.Request) {

	quaries := request.URL.Query()
	FirstNum := quaries.Get("first_num")
	SecondNum := quaries.Get("second_num")
	Operation := quaries.Get("operation")

	FirstNumInt := services.StringToInt(FirstNum)
	SecondNumInt := services.StringToInt(SecondNum)

	ActionOfCulc := services.Calc(FirstNumInt, SecondNumInt, Operation)
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
	services.WriteInJson(config)
}

func GetHistory(response http.ResponseWriter, request *http.Request) {

	method := request.Method
	if method != http.MethodGet {
		http.Error(response, "Method != Get", http.StatusBadRequest)
		return
	}
	services.RequestForHistory(response)

}

func CleanHistory(response http.ResponseWriter, request *http.Request) {
	services.Clean()
}
