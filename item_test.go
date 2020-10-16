package configure

import "testing"

var testItem = newItem(nil, "test", "hope everything right", "a simple test")

func TestItem_Description(t *testing.T) {
	r := testItem.Description()
	expected := "a simple test"
	if r != expected {
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}

func TestItem_Name(t *testing.T) {
	r := testItem.Name()
	expected := "test"
	if r != expected {
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}

func TestItem_Val(t *testing.T) {
	r := testItem.Val()
	expected := "hope everything right"
	if r != expected {
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}

