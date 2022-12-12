package services

import (
	"encoding/json"
	"io"
	"lesson15/internal/models"
	"lesson15/internal/reprository"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	Reprository *reprository.Reprository
}

func NewService(rep *reprository.Reprository) *Service {
	return &Service{Reprository: rep}
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

func (s *Service) WriteInJson(config models.CalcResult) {

	s.Reprository.InsertSmth(config)

	//err := r.InsertSmth(config)
	//
	//log.Println(err)

	//file, err := os.OpenFile("./Result.json", os.O_RDWR, 0777)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//contentJson, err := io.ReadAll(file)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//var History []models.CalcResult
	//err = json.Unmarshal(contentJson, &History)
	//if err != nil {
	//	log.Println(err)
	//}
	//History = append(History, config)
	//
	//bytes, err := json.MarshalIndent(History, "", "  ")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//err = os.WriteFile("Result.json", bytes, 0777)
	//if err != nil {
	//	log.Println(err)
	//}
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
		log.Println("asdasdasd", err)
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

//func CleanLast() {

//file, err := os.OpenFile("./Result.json", os.O_RDWR, 0777)
//if err != nil {
//	log.Println(err)
//	return
//}
//
//contentJson, err := io.ReadAll(file)
//if err != nil {
//	log.Println(err)
//}
//var History []models.CalcResult
//err = json.Unmarshal(contentJson, &History)
//if err != nil {
//	log.Println(err)
//}
//
//drob := contentJson[:len(contentJson)-1]
//
//delete, err := json.MarshalIndent(drob, "", "  ")
//if err != nil {
//	log.Println(err)
//	return
//}
//err = os.WriteFile("./Result.json", delete, 0777)
//if err != nil {
//	log.Println(err)
//	return
//}

//err = file.Truncate(0)
//file.Close()
//}

func (s *Service) Pagination(count, page int) ([]*models.Users, error) {
	return s.Reprository.Pagination(count, page)
}
