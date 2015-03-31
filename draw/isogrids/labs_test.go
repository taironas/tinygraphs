package isogrids

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDiagonals(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	Diagonals(rec, key, colorTheme[0], colorTheme[1], 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}

func TestHalfDiagonals(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	HalfDiagonals(rec, key, colorTheme[0], colorTheme[1], 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}

func TestSkeleton(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	Skeleton(rec, key, colorTheme[0], colorTheme[1], 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}
