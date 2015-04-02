package spaceinvaders

import "testing"

func TestnewInvader(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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

func TestArmsFromKey(t *testing.T) {
	t.Parallel()
	got := ArmsFromKey('c')
	expected := 2
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = ArmsFromKey('a')
	expected = 2
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestAnthenasFromKey(t *testing.T) {
	t.Parallel()
	got := AnthenasFromKey('d')
	expected := 0
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = AnthenasFromKey('a')
	expected = 1
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestLengthFromKey(t *testing.T) {
	t.Parallel()
	got := LengthFromKey('h')
	expected := 5
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = LengthFromKey('a')
	expected = 5
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestHeightFromKey(t *testing.T) {
	t.Parallel()
	got := HeightFromKey('j')
	expected := 6
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = HeightFromKey('a')
	expected = 7
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestEyesFromKey(t *testing.T) {
	t.Parallel()
	got := EyesFromKey('b')
	expected := 3
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = EyesFromKey('a')
	expected = 2
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestHasFootFromKey(t *testing.T) {
	t.Parallel()
	got := HasFootFromKey('a')
	expected := false
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
	got = HasFootFromKey('b')
	expected = true
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestHasArmsUpFromKey(t *testing.T) {
	t.Parallel()
	got := HasArmsUpFromKey('a')
	expected := false
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
	got = HasArmsUpFromKey('b')
	expected = true
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestArmSizeFromKey(t *testing.T) {
	t.Parallel()
	got := ArmSizeFromKey('a')
	expected := 3
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = ArmSizeFromKey('b')
	expected = 2
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestAnthenaSizeFromKey(t *testing.T) {
	t.Parallel()
	got := AnthenaSizeFromKey('a')
	expected := 2
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	got = AnthenaSizeFromKey('b')
	expected = 1
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}
