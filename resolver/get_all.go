package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cpunekar/graphql-go-pet/model"
	"io/ioutil"
	"net/http"
)

type GetAllResolver struct{}

// Get All Pets
func (r *GetAllResolver) GetPets(ctx context.Context) ([]*model.Pet, error) {
	url := baseURL + "/pets"
	response, err := http.Get(url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	pets, err := localGetPets(body)
	if err != nil {
		fmt.Println(err)
	}

	return pets, nil
}

// List of Pets
func localGetPets(body []byte) ([]*model.Pet, error) {
	pets := []*model.Pet{}
	err := json.Unmarshal(body, &pets)
	if err != nil {
		fmt.Println("Error getting pets")
	}
	return pets, err
}
