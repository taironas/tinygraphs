package spaceinvaders

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	colorTheme = []color.RGBA{
		{255, 245, 249, 255},
		{232, 70, 134, 255},
		{232, 70, 186, 255},
		{232, 70, 81, 255},
	}
	key  string
	keys []string
)

func init() {
	h := md5.New()
	io.WriteString(h, "hello")
	key = fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello2")
	key2 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello3")
	key3 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello4")
	key4 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello5")
	key5 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello6")
	key6 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello7")
	key7 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello8")
	key8 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello9")
	key9 := fmt.Sprintf("%x", h.Sum(nil)[:])

	keys = []string{key, key2, key3, key4, key5, key6, key7, key8, key9}

}

func TestSpaceInvaders(t *testing.T) {
	t.Parallel()
	for _, k := range keys {
		rec := httptest.NewRecorder()
		SpaceInvaders(rec, k, colorTheme, 100)
		if rec.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
		}
	}

}
