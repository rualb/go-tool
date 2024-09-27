package toolhttp

// TODO gzip-encoded-content
// TODO as http-client with config(gzip,baseurl,headers)

/*
!!! If you want to use generics at the struct level, which is more advanced but allows more flexibility, you would need to define the struct itself as generic:
*/
import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type StrMap = map[string]string

// EncodeURL encodes a string for safe inclusion in a URL query.
func EncodeURL(input string) string {
	return url.QueryEscape(input)

}

// URLEncode encodes a string for safe inclusion in a URL query.
func JoinURL(baseURL string, queryParams StrMap) (string, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %v", err)
	}

	if queryParams != nil {

		// Add query parameters from the map
		query := url.Values{}
		for key, value := range queryParams {
			query.Add(key, value)
		}
		parsedURL.RawQuery = query.Encode()
	}

	return parsedURL.String(), nil
}

func readBodyAllBytes(resp *http.Response) ([]byte, error) {

	defer resp.Body.Close()

	var reader io.Reader = resp.Body
	contentEncoding := resp.Header.Get("Content-Encoding")
	if contentEncoding != "" {

		switch {
		case strings.Contains(contentEncoding, "gzip"):
			// Handle gzip encoding
			gzipReader, err := gzip.NewReader(resp.Body)
			if err != nil {
				fmt.Println("Error creating gzip reader:", err)

			}
			defer gzipReader.Close()
			reader = gzipReader

		case strings.Contains(contentEncoding, "deflate"):
			// // Handle deflate encoding
			// reader = flate.NewReader(resp.Body)
			return nil, fmt.Errorf("deflate encoding is not supported")
		case strings.Contains(contentEncoding, "br"):
			// Handle Brotli encoding if needed
			// Brotli is not included in the Go standard library; use a third-party package
			// For example: github.com/andybalholm/brotli
			// Example:
			// brReader := brotli.NewReader(resp.Body)
			// reader = brReader
			return nil, fmt.Errorf("brotli encoding is not supported")
		default:

			return nil, fmt.Errorf("%s encoding is not supported", contentEncoding)

		}
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return data, err
}

func tryGetBytes(reqBody any) (bodyData []byte, contentType string, err error) {
	bodyData = nil
	contentType = ""

	if reqBody == nil {
		return bodyData, contentType, nil
	}

	switch v := reqBody.(type) {
	case string:
		bodyData = []byte(v)
		contentType = "text/plain"
	case []byte:
		bodyData = v
		contentType = "application/octet-stream"
	default:
		// Attempt to JSON-encode the body
		bodyData, err = json.Marshal(v)
		if err != nil {
			return nil, "", fmt.Errorf("error encoding to JSON: %v", err)
		}
		contentType = "application/json"
	}
	return bodyData, contentType, nil

}

// HttpGetBinary makes an HTTP GET request to the specified opt.BaseURL with query parameters from opt.QueryParams
// and returns the binary data from the response.
func GetBytes(baseURL string, queryParams StrMap, reqBody any) ([]byte, error) {

	var resp *http.Response
	// G107: Potential HTTP request made with variable url
	fullURL, err := JoinURL(baseURL, queryParams)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}

	parsedURL, err := url.Parse(fullURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid URL: %s", fullURL)
	}

	fullURL = parsedURL.String()

	bodyData, contentType, err := tryGetBytes(reqBody)
	if err != nil {
		return nil, err
	}

	// Potential HTTP request made with variable url
	if bodyData != nil {
		// Make the POST request
		//nolint:G107
		resp, err = http.Post(fullURL, contentType, bytes.NewBuffer(bodyData)) //nolint:G107
		if err != nil {
			return nil, fmt.Errorf("failed to make GET request: %v", err)
		}
	} else {
		// Make the GET request
		//nolint:G107
		resp, err = http.Get(fullURL) //nolint:G107
		if err != nil {
			return nil, fmt.Errorf("failed to make GET request: %v", err)
		}
	}

	// Read the binary data from the response
	data, err := readBodyAllBytes(resp)
	if err != nil {
		return nil, err
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	return data, nil
}

func GetJSON[T any](baseURL string, queryParams StrMap, reqBody any) (T, error) {

	data, err := GetBytes(baseURL, queryParams, reqBody)

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

func GetText(baseURL string, queryParams StrMap, reqBody any) (string, error) {

	data, err := GetBytes(baseURL, queryParams, reqBody)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
