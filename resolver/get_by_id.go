package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cpunekar/graphql-go-pet/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

type GetByIdResolver struct{}

// Get Pet by Id
func (r *GetByIdResolver) GetPetById(ctx context.Context, args struct{ ID int32 }) ([]*model.Pet, error) {
	url := baseURL + "/pets/" + strconv.FormatInt(int64(args.ID), 10)
	fmt.Println(url)
	response, err := http.Get(url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	pet, err := localGetPet(body)
	if err != nil {
		fmt.Println(err)
	}

	return pet, nil
}

// Single Pet by Id
func localGetPet(body []byte) ([]*model.Pet, error) {
	pets := []*model.Pet{}
	err := json.Unmarshal(body, &pets)
	if err != nil {
		fmt.Println("Error getting pets")
	}
	return pets, err
}
