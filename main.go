package main

import (
	"encoding/json"
	"net/http"
)

const (
	breedsEndpoint   = "https://dog.ceo/api/breeds/list/all"
	breedImageFormat = "https://dog.ceo/api/breed/%s/images/random"
)

type BreedsResponse struct {
	Message map[string][]string `json:"message"`
}

func main() {
	breeds, err := getAllBreeds()
}

func getAllBreeds() (map[string][]string, error) {
	resp, err := http.Get(breedsEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var breedsResponse BreedsResponse
	err = json.NewDecoder(resp.Body).Decode(&breedsResponse)
	if err != nil {
		return nil, err
	}

	return breedsResponse.Message, nil
}
