package sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	Credentials Credentials
	Url         string
	Aggregate   bool
}

func jsonEncode(data map[string]string) (*bytes.Buffer, error) {
	jsonStr, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer([]byte(jsonStr)), nil
}

func parseResponse(content io.ReadCloser) (Response, error) {
	dec := json.NewDecoder(content)
	json := make(map[string]interface{})
	err := dec.Decode(&json)

	if err != nil {
		return Response{}, err
	}

	return Response{Data: json}, nil
}

func (c Client) Request(endpoint string, data map[string]string) (Response, error) {
	data["clientId"] = c.Credentials.ClientId
	if c.Aggregate {
		data["aggregateResponse"] = "1"
	}

	jsonData, err := jsonEncode(data)

	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", c.Url+endpoint, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Credentials.Bearer)

	if err != nil {
		return Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return Response{}, err
	}

	response, err := parseResponse(resp.Body)

	if err != nil {
		return Response{}, err
	}

	return response, nil
}
