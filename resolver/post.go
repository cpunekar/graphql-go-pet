package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cpunekar/graphql-go-pet/model"
	"net/http"
	"strings"
)

type PostResolver struct {
}

// Create Pet
func (r *PostResolver) AddPet(ctx context.Context, args struct{ Input *model.Pet }) (*model.Pet, error) {
	url := baseURL + "/pets"

	p := model.Pet{
		Name: args.Input.Name,
		Tag:  args.Input.Tag,
	}

	payload, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("Calling URL %v to add pet with payload %v \n", url, string(payload))

	payloadStr := string(payload)
	response, err := postDataRequest(url, &payloadStr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Response: %v", response)

	return &p, nil
}

func postDataRequest(url string, payload *string) (*http.Response, error) {
	client := &http.Client{}
	var req *http.Request
	if payload == nil {
		req, _ = http.NewRequest(http.MethodPost, url, nil)
	} else {
		req, _ = http.NewRequest(http.MethodPost, url, strings.NewReader(*payload))
	}

	fmt.Printf("Request %v", req)
	resp, err := client.Do(req)

	return resp, err
}
