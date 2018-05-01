package callers

import (
	"encoding/json"
	"errors"
	"reflect"
	. "leo/models/caller_options"
	. "leo/models/queries"
	. "leo/models/responses"
	"libs/go.uuid"
)

type DialogFlowClient struct {
	accessToken string
	apiLang     string
	apiVersion  string
	apiBaseUrl  string
	sessionID   string
}

// Creates a DialogFlowClient instance. If none AgentClientOptions are provided uses default values
func NewDialogFlowClient(options AgentClientOptions) (error, *DialogFlowClient) {
	client := &DialogFlowClient{}

	client.accessToken = options.AccessToken
	if client.accessToken == "" {
		client.accessToken = API_KEY
	}

	client.apiBaseUrl = options.ApiBaseUrl
	if client.apiBaseUrl == "" {
		client.apiBaseUrl = DEFAULT_BASE_URL
	}

	client.apiLang = options.ApiLang
	if client.apiLang == "" {
		client.apiLang = DEFAULT_CLIENT_LANG
	}

	client.apiVersion = options.ApiVersion
	if client.apiVersion == "" {
		client.apiVersion = DEFAULT_API_VERSION
	}

	client.sessionID = options.SessionID
	if client.sessionID == "" {
		u, _ := uuid.NewV4()
		client.sessionID = u.String()
	}

	return nil, client
}

// Initialize a new HTTP request
func NewDialogFlowRequest(client *DialogFlowClient, overridedRequestOptions RequestOptions) *Request {
	headers := map[string]string{
		"Authorization": "Bearer " + client.GetAgentAccessToken(),
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

//Performs an http request using <this> DialogFlowClient and provided AgentQuery struct,
//returns response data along with any errors.
func (client *DialogFlowClient) AgentQueryCreateRequest(query AgentQuery) (AgentQueryResponse, error) {
	var response AgentQueryResponse

	if reflect.DeepEqual(query, AgentQuery{}) {
		return response, errors.New("query cannot be empty")
	}

	request := NewDialogFlowRequest(
		client,
		RequestOptions{
			URI:    client.GetAgentBaseUrl() + "query?v=" + client.GetAgentApiVersion(),
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

// GET DialogFlow access token
func (client *DialogFlowClient) GetAgentAccessToken() string {
	return client.accessToken
}

// GET dialog_flow version
func (client *DialogFlowClient) GetAgentApiVersion() string {
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
func (client *DialogFlowClient) GetAgentBaseUrl() string {
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
