package server

import (
	"encoding/json"
	"io"
	"lesson15/internal/models"
	"lesson15/internal/services"
	"log"
	"net/http"
	"os"
)

func (s *Server) Calculation(response http.ResponseWriter, request *http.Request) {

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
	s.Service.WriteInJson(config)
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

func CleanLast(response http.ResponseWriter, request *http.Request) {
	var History []models.CalcResult

	file, err := os.OpenFile("Result.json", os.O_RDWR, 0777)
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
		return
	}
	//log.Println(string(contentJson))

	drob := History[:len(History)-1]

	deleted, err := json.MarshalIndent(drob, "", "  ")
	//deleted, err := json.Marshal(drob)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.WriteFile("./Result.json", deleted, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	//services.CleanLast()
}

func (s *Server) Pagination(response http.ResponseWriter, request *http.Request) {
	quaries := request.URL.Query()
	Count := quaries.Get("count")
	Page := quaries.Get("page")

	FirstNumInt := services.StringToInt(Count)
	SecondNumInt := services.StringToInt(Page)

	pagination, err := s.Service.Pagination(FirstNumInt, SecondNumInt)
	if err != nil {
		log.Println(err)
		return
	}

	//for _, getUser := range Folders {
	//	log.Println("*", getUser)

	data, err := json.MarshalIndent(pagination, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))
	_, err = response.Write(data)
	if err != nil {
		log.Println(err)
		return
	}

}
