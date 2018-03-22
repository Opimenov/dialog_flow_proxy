package callers

import (
	"encoding/json"
	"errors"
	"reflect"
	. "leo/models/caller_options"
	. "leo/models/queries"
	. "leo/models/responses"
)

type EngineeringClient struct {
	accessToken string
	apiBaseUrl  string
}

// Create API.AI instance
func NewEngineeringClient(options EngineeringClientOptions) (error, *EngineeringClient) {
	client := &EngineeringClient{}
	client.accessToken = options.AccessToken
	if client.accessToken == "" {
		client.accessToken = API_KEY
	}
	client.apiBaseUrl = options.ApiBaseUrl
	if client.apiBaseUrl == "" {
		client.apiBaseUrl = DEFAULT_BASE_URL
	}
	return nil, client
}

// Initialize a new HTTP request
func NewDialogFlowRequest(client *DialogFlowClient, overridedRequestOptions RequestOptions) *Request {
	headers := map[string]string{
		"Authorization": "Bearer " + client.GetAccessToken(),
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

// Takes natural language text and information as query parameters and returns information as JSON
func (client *DialogFlowClient) QueryFindRequest(query AgentQuery) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(query, AgentQuery{}) {
		return response, errors.New("query cannot be empty")
	}

	if query.V == "" {
		query.V = client.GetApiVersion()
	}

	if query.Lang == "" {
		query.Lang = client.GetApiLang()
	}

	if query.SessionID == "" {
		query.SessionID = client.GetSessionID()
	}

	request := NewDialogFlowRequest(
		client,
		RequestOptions{
			URI:         client.GetBaseUrl() + "query",
			Method:      "GET",
			Body:        nil,
			QueryParams: query.ToMap(),
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Takes natural language text and information as JSON in the POST body and returns information as JSON
func (client *DialogFlowClient) QueryCreateRequest(query AgentQuery) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(query, AgentQuery{}) {
		return response, errors.New("query cannot be empty")
	}

	request := NewDialogFlowRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "query?v=" + client.GetApiVersion(),
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

// GET API.AI access token
func (client *DialogFlowClient) GetAccessToken() string {
	return client.accessToken
}

// GET dialog_flow version
func (client *DialogFlowClient) GetApiVersion() string {
	if client.apiVersion != "" {
		return client.apiVersion
	}
	return DEFAULT_API_VERSION
}

// GET api language
func (client *DialogFlowClient) GetApiLang() string {
	if client.apiLang != "" {
		return client.apiLang
	}
	return DEFAULT_CLIENT_LANG
}

// Get api base url
func (client *DialogFlowClient) GetBaseUrl() string {
	if client.apiBaseUrl != "" {
		return client.apiBaseUrl
	}
	return DEFAULT_BASE_URL
}

// Get current session ID
func (client *DialogFlowClient) GetSessionID() string {
	return client.sessionID
}

// Set a new seesion ID
func (client *DialogFlowClient) SetSessionID(sessionID string) {
	client.sessionID = sessionID
}
