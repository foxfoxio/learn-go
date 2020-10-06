package hellogo

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHellpGopher(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	HelloGopher(rr, req)
	if r := rr.Result(); r.StatusCode != http.StatusOK {
		t.Errorf("Gopher StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
	} else {
		responseData, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "Hello, Gopher.\n", string(responseData))
	}
}
