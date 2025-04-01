package contribution

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetContribution_FixedContribution(t *testing.T) {
	// Set Gin to test mode, initialize a router, and register the POST /contribution route with the SetContribution handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/contribution", SetContribution)

	// Create a mock POST request with a JSON body and a response recorder to capture the handler's output
	body := `{"contribution": 100.0}`
	req, _ := http.NewRequest(http.MethodPost, "/contribution", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Simulate sending the mock request to the router and capture the response in the recorder
	router.ServeHTTP(rec, req)

	// Check if the response status code is 200 OK; if not, log an error with the expected and actual status codes
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	// Parse the JSON response body into a map; if parsing fails, log the error and stop the test
	var response map[string]float64
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Define the expected response and compare it with the actual response; log an error if they don't match
	expected := map[string]float64{"Fixed contribution": 100.0}
	if response["Fixed contribution"] != expected["Fixed contribution"] {
		t.Errorf("Expected response %v, but got %v", expected, response)
	}
}
