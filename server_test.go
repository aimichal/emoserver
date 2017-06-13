package main

import "testing"
import "os"
import "net/http"
import "net/http/httptest"

func TestServerDescriptionHappy(t *testing.T) {
	os.Setenv("HOSTNAME", "abc123")
	happy = true

	if serverDescription() != "Very Happy server on host abc123" {
		t.Fatal("Test failed, yo")
	}
}

func TestServerDescriptionSad(t *testing.T) {
	os.Setenv("HOSTNAME", "abc123")
	happy = false

	if serverDescription() != "Very Sad server on host abc123" {
		t.Fatal("Test failed, yo")
	}
}

func confirm(t *testing.T, handler http.HandlerFunc, path string, expectedStatus int, expectedBody string) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}

}

func TestHappy(t *testing.T) {
	os.Setenv("HOSTNAME", "abc123")
	happy = false
	confirm(t, MakeHappy, "/something", 200, "Very Sad server on host abc123 is now happy\n")
	confirm(t, Something, "/something", 200, "Very Happy server on host abc123 handling request: /something\n")
}

func TestSad(t *testing.T) {
	os.Setenv("HOSTNAME", "abc123")
	happy = true
	confirm(t, MakeSad, "/something", 200, "Very Happy server on host abc123 is now sad\n")
	confirm(t, Something, "/something", 500, "Very Sad server on host abc123 handling request: /something\n")
}
