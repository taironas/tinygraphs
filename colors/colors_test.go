package colors

import "testing"

func TestMapOfColorsThemes(t *testing.T) {
	m := MapOfColorThemes()

	// can access base color
	if base, ok := m["base"]; !ok {
		t.Errorf("base color expected")
	} else {
		if len(base) != 2 {
			t.Errorf("expects", 2, "got", len(base))
		}
	}

	// can access all themes:
	themes := []string{
		"sugarsweets",
		"heatwave",
		"daisygarden",
		"seascape",
		"summerwarmth",
		"bythepool",
		"duskfalling",
		"frogideas",
		"berrypie",
	}

	for _, v := range themes {
		if got, ok := m[v]; !ok {
			t.Errorf("color", v, "expected")
		} else {
			if len(got) != 4 {
				t.Errorf("expects", 4, "got", len(got))
			}
		}
	}
}
