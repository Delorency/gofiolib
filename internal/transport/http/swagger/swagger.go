package swagger

import (
	"time"
)

type SwaggerPerson struct {
	ID         uint      `json:"Id"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
	DeletedAt  time.Time `json:"DeletedAt"`
	Name       string    `json:"Name"`
	Surname    string    `json:"Surname"`
	Patronymic string    `json:"Patronymic"`
	Age        uint8     `json:"Age"`
	Gender     string    `json:"Gender"`
	Nat        string    `json:"Nat"`
}
type SwaggerValidateData struct {
	Err    string   `json:"error"`
	Fields []string `json:"fields"`
}
type SwaggerNewError struct {
	Err string `json:"error"`
}
