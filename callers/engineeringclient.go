//Contains structs and functionality to perform DialogFlow.com and
// Engineering.com api calls
package callers

import (
	"encoding/json"
	"errors"
	"reflect"
	. "leo/models/caller_options"
	. "leo/models/queries"
	. "leo/models/responses"
	"fmt"
)

//Contains necessary data for making Engineering.com api calls.
type EngineeringClient struct {
	AccessToken string
	ApiBaseUrl  string
}

//Creates new Engineering.com client using provided options,
// or default values if none provided.
func NewEngineeringClient(options EngineeringClientOptions) (error, *EngineeringClient) {
	client := &EngineeringClient{}
	client.AccessToken = options.AccessToken
	if client.AccessToken == "" {
		client.AccessToken = API_KEY
	}
	client.ApiBaseUrl = options.ApiBaseUrl
	if client.ApiBaseUrl == "" {
		client.ApiBaseUrl = DEFAULT_BASE_URL
	}
	return nil, client
}

// Initialize a new HTTP request
func NewEngineeringRequest(client *EngineeringClient, overridedRequestOptions RequestOptions) *Request {
	headers := map[string]string{
		"Authorization": "Bearer " + client.GetEngineeringAccessToken(),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}

	request := &Request{
		URI:         overridedRequestOptions.URI,
		Method:      overridedRequestOptions.Method,
		Headers:     headers,
		QueryParams: overridedRequestOptions.QueryParams,
		Body:        overridedRequestOptions.Body,
	}

	return request
}

//Makes a request and returns a result in a form of an AgentQueryResponse.
// Uses POST method.
func (client *EngineeringClient) QueryCreateRequest(query EngineeringQuery) (AgentQueryResponse, error) {
	var response AgentQueryResponse
	if reflect.DeepEqual(query, EngineeringQuery{}) {
		return response, errors.New("query cannot be empty")
	}

	request := NewEngineeringRequest(
		client,
		RequestOptions{
			URI:    client.GetEngineeringBaseUrl() + "query",
			Method: "POST",
			Body:   query,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func (client *EngineeringClient) GetEngineeringAccessToken() string {
	return client.AccessToken
}

func (client *EngineeringClient) GetEngineeringBaseUrl() string {
	if client.ApiBaseUrl != "" {
		return client.ApiBaseUrl
	}
	return DEFAULT_BASE_URL
}
//Should make and API call to Engineering.com to create a project.
//All the necessary pieces are ready, just need to put it together.
func CreateProject(proj_name string) (answer EngineeringQueryResponse) {
	//TODO provided that there is engineering.com API access token available implement object creation.
	//Infrastructure is in place just put it together.
	fmt.Println("creating project with name " + proj_name)
	return
}
