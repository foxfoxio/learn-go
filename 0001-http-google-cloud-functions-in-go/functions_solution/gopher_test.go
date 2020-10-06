package hellogo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGopher(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	Gopher(rr, req)
	if r := rr.Result(); r.StatusCode != http.StatusOK {
		t.Errorf("Gopher StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
	} else {
		contentType := r.Header.Get("Content-Type")
		assert.Equal(t, "image/png", contentType)
	}
}
