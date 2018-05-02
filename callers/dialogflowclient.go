//Contains structs and functionality to perform DialogFlow.com and
//Engineering.com api calls
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

//Defines necessary fields for making DialogFlow api calls.
type DialogFlowClient struct {
	AccessToken string
	ApiLang     string
	ApiVersion  string
	ApiBaseUrl  string
	SessionID   string
}

// Creates a DialogFlowClient instance.
// If none AgentClientOptions are provided uses default values
func NewDialogFlowClient(options AgentClientOptions) (error, *DialogFlowClient) {
	client := &DialogFlowClient{}

	client.AccessToken = options.AccessToken
	if client.AccessToken == "" {
		client.AccessToken = API_KEY
	}

	client.ApiBaseUrl = options.ApiBaseUrl
	if client.ApiBaseUrl == "" {
		client.ApiBaseUrl = DEFAULT_BASE_URL
	}

	client.ApiLang = options.ApiLang
	if client.ApiLang == "" {
		client.ApiLang = DEFAULT_CLIENT_LANG
	}

	client.ApiVersion = options.ApiVersion
	if client.ApiVersion == "" {
		client.ApiVersion = DEFAULT_API_VERSION
	}

	client.SessionID = options.SessionID
	if client.SessionID == "" {
		u, _ := uuid.NewV4()
		client.SessionID = u.String()
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

//Performs an http request using <this> DialogFlowClient
// and provided AgentQuery struct,
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

//Returns DialogFlow access token
func (client *DialogFlowClient) GetAgentAccessToken() string {
	return client.AccessToken
}

//Returns dialog_flow version
func (client *DialogFlowClient) GetAgentApiVersion() string {
	if client.ApiVersion != "" {
		return client.ApiVersion
	}
	return DEFAULT_API_VERSION
}

//Returns api language
func (client *DialogFlowClient) GetApiLang() string {
	if client.ApiLang != "" {
		return client.ApiLang
	}
	return DEFAULT_CLIENT_LANG
}

//Returns base url
func (client *DialogFlowClient) GetAgentBaseUrl() string {
	if client.ApiBaseUrl != "" {
		return client.ApiBaseUrl
	}
	return DEFAULT_BASE_URL
}

//Returns current session ID
func (client *DialogFlowClient) GetSessionID() string {
	return client.SessionID
}

// Sets a new seesion ID
func (client *DialogFlowClient) SetSessionID(sessionID string) {
	client.SessionID = sessionID
}
