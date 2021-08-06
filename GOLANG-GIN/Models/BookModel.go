package Models

import (
	//"github.com/thedevsaddam/govalidator"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Book struct {
	Id           uint      `validate:"omitempty,uuid" json:"id"`
	Title        string    `validate:"required" json:"title"`
	Author       string    `validate:"required" json:"author"`
	Description  string    `validate:"required" json:"description"`
	Price        float64   `validate:"required" json:"price"`
	Isbn         string    `validate:"required" json:"isbn"`
	Release_date time.Time `validate:"isdafault" json:"release_date"`
}

func main() {
	book := Book{
		Id:          1,
		Title:       "morning",
		Author:      "James",
		Description: "behalf",
		Price:       200.0,
		Isbn:        "54638",
	}
	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (b *Book) TableName() string {
	return "book"
}
