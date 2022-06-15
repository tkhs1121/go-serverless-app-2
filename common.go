package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func parseRequest(payload Payload) (string, error) {
	var req Request
	err := json.NewDecoder(strings.NewReader(payload.Body)).Decode(&req)
	if err != nil {
		return "", fmt.Errorf("cannot decode payload: %v: %v", err, payload.Body)
	}
	return req.URL, nil
}

func checkAmazonURL(url string) (bool, error) {
	return regexp.MatchString(`^http(s)?://(www.)?amazon.(co.)?jp[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]*`, url)
}
