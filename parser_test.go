package configure

import "testing"

var parser1 = newParser("# test\nkey1 = 123\nkey2 = 1234adf\n[sec1]\nkey3 = aad\nkey4 = aad",  newOption())
var parser2 = newParser("[sec0]\n# test\nkey1 = 123\nkey2 = 1234adf\n[sec1]\nkey3 = aad\nkey4 = aad",  newOption())

func TestParser_ParseName(t *testing.T) {
	var f *File
	f = newFile(nil, "")
	parser1.parse(f)

	defaultName := f.sections[0].Name()
	expected := "default"

	if expected != defaultName {
		t.Errorf("expected '%s' but got '%s'", expected, defaultName)
	}
}

func TestParser_ParseNameNoDefault(t *testing.T) {
	var f *File
	f = newFile(nil, "")
	parser2.parse(f)

	defaultName := f.sections[0].Name()
	expected := "sec0"

	if expected != defaultName {
		t.Errorf("expected '%s' but got '%s'", expected, defaultName)
	}
}


