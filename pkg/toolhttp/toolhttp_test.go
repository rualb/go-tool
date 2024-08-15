package toolhttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

var (
	server     *httptest.Server
	serverOnce sync.Once
)

func startTestServer() {
	serverOnce.Do(func() {
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			switch r.URL.Path {
			case "/bytes":
				w.Write([]byte(q.Get("message")))
			case "/json":

				json.NewEncoder(w).Encode(map[string]string{"message": q.Get("message")})

			case "/json_post":
				if r.Method == "POST" {

					x := StrMap{}
					if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
						panic(err)
					}
					json.NewEncoder(w).Encode(&map[string]string{"message": x["message"]})

				}
			default:
				http.NotFound(w, r)
			}
		}))
	})
}

func stopTestServer() {
	if server != nil {
		server.Close()
	}
}

func TestMain(m *testing.M) {
	// Start the test server
	startTestServer()

	// Run the tests
	code := m.Run()

	// Stop the test server
	stopTestServer()
	// Exit with the code returned by m.Run()
	os.Exit(code)
}

func TestGetBytes(t *testing.T) {

	message := "Hello, World!"
	args := StrMap{
		"message": message,
	}

	httpClient := HttpClient{server.URL + "/bytes", args, ""}
	// Call the function
	data, err := GetBytes(&httpClient, nil)
	if err != nil {
		t.Fatalf("GetBytes() error = %v", err)
	}

	// Check the result
	expected := []byte(message)
	if string(data) != string(expected) {
		t.Errorf("GetBytes() = %v, want %v", string(data), string(expected))
	}
}

func TestGetJSON(t *testing.T) {

	message := "Hello, World!"
	args := StrMap{
		"message": message,
	}

	// Define the expected result type
	type resultType struct {
		Message string `json:"message"`
	}
	httpClient := HttpClient{server.URL + "/json", args, ""}
	// Call the function
	var result resultType
	result, err := GetJSON[resultType](&httpClient, nil)
	if err != nil {
		t.Fatalf("GetJSON() error = %v", err)
	}

	// Check the result
	expected := message
	if result.Message != expected {
		t.Errorf("GetJSON() = %v, want %v", result.Message, expected)
	}
}
func TestGetJSON_POST(t *testing.T) {

	message := "Hello, World!"

	// Define the expected result type
	type resultType struct {
		Message string `json:"message"`
	}
	httpClient := HttpClient{server.URL + "/json_post", nil, ""}
	// Call the function
	var result resultType
	result, err := GetJSON[resultType](&httpClient, &StrMap{"message": message})
	if err != nil {
		t.Fatalf("GetJSON() error = %v", err)
	}

	// Check the result
	expected := message
	if result.Message != expected {
		t.Errorf("GetJSON() = %v, want %v", result.Message, expected)
	}
}

func TestGetText(t *testing.T) {

	message := "Hello, World!"
	args := StrMap{
		"message": message,
	}

	httpClient := HttpClient{server.URL + "/bytes", args, ""}
	// Call the function
	result, err := GetText(&httpClient, nil)
	if err != nil {
		t.Fatalf("GetText() error = %v", err)
	}

	// Check the result
	expected := message
	if result != expected {
		t.Errorf("GetText() = %v, want %v", result, expected)
	}
}
