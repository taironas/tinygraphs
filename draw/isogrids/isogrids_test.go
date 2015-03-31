package isogrids

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsogrids(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	Isogrids(rec, key, colorTheme, 10, 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

}
