package configure

import "testing"

func TestUtils_deleteEmpty(t *testing.T) {
	text := []string{"aaa", "", "dfasdfa", "", "adfadfadf"}
	r := len(deleteEmpty(text))
	expected := 3
	if r != expected {
		t.Errorf("expected '%d' but got '%d'", expected, r)
	}
}

func TestUtils_getSecName(t *testing.T) {
	text := "[aaaadfweq]"
	r := getSecName(text)
	expected := "aaaadfweq"
	if r != expected {
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}

func TestUtils_splitFirst(t *testing.T) {
	text := "# aaaadfweqdfqweq132412dfbvv"
	r := splitFirst(text)
	expected := "aaaadfweqdfqweq132412dfbvv"
	if r != expected {
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}
