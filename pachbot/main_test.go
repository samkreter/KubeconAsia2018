// endpoints_test.go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestACRPushHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	b, err := ioutil.ReadFile("testfiles/acrExampleResp.json")
	if err != nil {
		assert.NoError(t, err, "Should not get error")
		t.Fail()
	}

	req, err := http.NewRequest("POST", "/acrpush", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(acrPushHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Printf("body: %s", rr.Body.String())
	// Check the response body is what we expect.
	// expected := `{"alive": true}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
