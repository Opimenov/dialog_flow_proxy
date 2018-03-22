package callers

import (
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Request struct {
	URI         string
	Method      string
	Headers     map[string]string
	Body        interface{}
	QueryParams map[string]string
}

func (r *Request) Perform() ([]byte, error) {
	var data []byte
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URI, nil)

	if r.Method != "GET" {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(r.Body)
		req, err = http.NewRequest(r.Method, r.URI, b)
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range r.QueryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	if err != nil {
		return data, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
