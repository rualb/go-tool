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

func TestReadAllBytes(t *testing.T) {

	message := "Hello, World!"
	args := str2str{
		"message": message,
	}

	// Call the function
	data, err := ReadAllBytes(server.URL+"/bytes", args)
	if err != nil {
		t.Fatalf("ReadAllBytes() error = %v", err)
	}

	// Check the result
	expected := []byte(message)
	if string(data) != string(expected) {
		t.Errorf("ReadAllBytes() = %v, want %v", string(data), string(expected))
	}
}

func TestReadAllJSON(t *testing.T) {

	message := "Hello, World!"
	args := str2str{
		"message": message,
	}

	// Define the expected result type
	type resultType struct {
		Message string `json:"message"`
	}

	// Call the function
	var result resultType
	result, err := ReadAllJSON[resultType](server.URL+"/json", args)
	if err != nil {
		t.Fatalf("ReadAllJSON() error = %v", err)
	}

	// Check the result
	expected := message
	if result.Message != expected {
		t.Errorf("ReadAllJSON() = %v, want %v", result.Message, expected)
	}
}

func TestReadAllText(t *testing.T) {

	message := "Hello, World!"
	args := str2str{
		"message": message,
	}

	// Call the function
	result, err := ReadAllText(server.URL+"/bytes", args)
	if err != nil {
		t.Fatalf("ReadAllText() error = %v", err)
	}

	// Check the result
	expected := message
	if result != expected {
		t.Errorf("ReadAllText() = %v, want %v", result, expected)
	}
}
