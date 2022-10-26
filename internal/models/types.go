package models

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type CalcResult struct {
	FirstNum  int    `json:"firstNum"`
	SecondNum int    `json:"secondNum"`
	Operation string `json:"operation"`
	Result    int    `json:"result"`
}
type HistoryElement struct {
	HistoryOfElements []CalcResult
}
