package personservice

import (
	"encoding/json"
	"fiolib/internal/models"
	"fmt"
	"log"
	"net/http"
)

func (ps *personService) Create(data *models.Person) error {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", data.Name))
	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %v", err)
	}

	resp, err = http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", data.Name))
	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %v", err)
	}

	resp, err = http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", data.Name))
	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}
	type CountryInfo struct {
		CountryID string `json:"country_id"`
	}
	type Response struct {
		Countries []CountryInfo `json:"country"`
	}
	var r Response
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %v", err)
	}
	if len(r.Countries) > 0 {
		data.Nat = r.Countries[0].CountryID
	} else {
		data.Nat = ""
	}

	return ps.repo.Create(data)
}
