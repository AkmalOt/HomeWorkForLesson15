package models

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type CalcResult struct {
	FirstNum  int    `gorm:"first_num"`
	SecondNum int    `gorm:"second_number"`
	Operation string `gorm:"operation"`
	Result    int    `gorm:"result"`
}

type Users struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone int    `json:"phone"`
}
