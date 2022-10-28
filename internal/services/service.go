package services

import (
	"encoding/json"
	"io"
	"lesson15/internal/models"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// -----------------------------------------------------------------------------

// For Calculate

func StringToInt(FirstNum string) int {
	FirstNumInt, err := strconv.Atoi(FirstNum)
	if err != nil {
		log.Println(err)
		return 0
	}

	return FirstNumInt
}
func Calc(FirstNumInt int, SecondNumInt int, Operation string) int {
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
	return ActionOfCulc
}

func WriteInJson(config models.CalcResult) {
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

// -----------------------------------------------------------------------------

// For GetHistory

func RequestForHistory(response http.ResponseWriter) {

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

//------------------------------------------------------

// for CleanHistory

func Clean() {
	file, err := os.OpenFile("./Result.json", os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return
	}

	err = file.Truncate(0)
	file.Close()
}
