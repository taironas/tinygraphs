package spaceinvaders

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
)

var (
	key string
)

func init() {
	h := md5.New()
	io.WriteString(h, "hello")
	key = fmt.Sprintf("%x", h.Sum(nil)[:])
}

func TestnewInvader(t *testing.T) {

	invs := []invader{newInvader(key), newInvader(key)}

	for _, inv := range invs {
		if inv.legs != 4 {
			t.Errorf("expected %d got %d", 4, inv.legs)
		}

		if inv.foot != false {
			t.Errorf("expected %v got %v", false, inv.legs)
		}

		if inv.arms != 2 {
			t.Errorf("expected %v got %v", 2, inv.arms)
		}

		if inv.anthenas != 0 {
			t.Errorf("expected %v got %v", 0, inv.anthenas)
		}

		if inv.length != 10 {
			t.Errorf("expected %v got %v", 10, inv.length)
		}

		if inv.height != 9 {
			t.Errorf("expected %v got %v", 9, inv.height)
		}

		if inv.eyes != 3 {
			t.Errorf("expected %v got %v", 3, inv.eyes)
		}

		if inv.armsUp != true {
			t.Errorf("expected %v got %v", true, inv.armsUp)
		}

		if inv.armSize != 3 {
			t.Errorf("expected %v got %v", 3, inv.armSize)
		}

		if inv.anthenaSize != 0 {
			t.Errorf("expected %v got %v", 0, inv.anthenaSize)
		}
	}
}

func TestLegsFromKey(t *testing.T) {
	got := LegsFromKey('c')
	expected := 2
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = LegsFromKey('a')
	expected = 3
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}

}
