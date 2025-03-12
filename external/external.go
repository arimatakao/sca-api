package external

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type Breed struct {
	Name     string `json:"name"`
	AltNames string `json:"alt_names,omitempty"`
}

func IsBreedAllowed(breedName string) (bool, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var breeds []Breed
	if err := json.Unmarshal(body, &breeds); err != nil {
		return false, err
	}

	for _, breed := range breeds {
		if breed.Name == breedName {
			return true, nil
		}
		log.Println(breed.AltNames)
		if breed.AltNames != "" {
			splitedAltNames := strings.Split(breed.AltNames, ", ")
			for _, altName := range splitedAltNames {
				if altName == breedName {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
