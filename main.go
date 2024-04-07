package main

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
