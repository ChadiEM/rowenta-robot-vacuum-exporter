package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseUrl[T any](endpoint string) (T, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return *new(T), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *new(T), err
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return *new(T), err
	}

	return result, nil
}
