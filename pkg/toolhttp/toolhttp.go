package toolhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type str2str = map[string]string

// URLEncode encodes a string for safe inclusion in a URL query.
func URLEncode(input string) string {
	return url.QueryEscape(input)
}

// HttpGetBinary makes an HTTP GET request to the specified baseURL with query parameters from args
// and returns the binary data from the response.
func ReadAllBytes(baseURL string, args str2str) ([]byte, error) {
	// Parse the base URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %v", err)
	}

	// Add query parameters from the map
	query := url.Values{}
	for key, value := range args {
		query.Add(key, value)
	}
	parsedURL.RawQuery = query.Encode()

	// Make the GET request
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the binary data from the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	return data, nil
}

func ReadAllJSON[T any](baseURL string, args str2str) (T, error) {

	data, err := ReadAllBytes(baseURL, args)

	res := new(T)

	if err != nil {

		return *res, err
	}

	// if err := json.NewDecoder(bytes.NewReader(data)).Decode(res); err != nil {
	// 	return *res, fmt.Errorf("error to decode JSON response: %v", err)
	// }

	if err := json.Unmarshal(data, res); err != nil {
		return *res, fmt.Errorf("error decoding JSON: %v", err)
	}

	return *res, nil
}

func ReadAllText(baseURL string, args str2str) (string, error) {

	data, err := ReadAllBytes(baseURL, args)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
