package validator_breed

import (
	"encoding/json"
	"net/http"
)

const url = "https://api.thecatapi.com/v1/breeds"

type Breed struct {
	Name string `json:"name"`
}

func ValidateBreed(breed string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var breeds []Breed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		return false
	}

	for _, b := range breeds {
		if b.Name == breed {
			return true
		}
	}
	return false
}
