package write

import (
	"image"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImageWithTemplate(t *testing.T) {
	t.Parallel()
	recorder := httptest.NewRecorder()
	var img image.Image = image.NewRGBA(image.Rect(0, 0, 0, 0))

	ImageWithTemplate(recorder, &img)
	if recorder.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
	}
}

func TestImageJPEG(t *testing.T) {
	t.Parallel()
	recorder := httptest.NewRecorder()
	var img image.Image = image.NewRGBA(image.Rect(0, 0, 0, 0))

	ImageJPEG(recorder, &img)
	if recorder.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
	}
}
