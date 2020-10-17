package configure

import (
	"testing"
)

func TestFile_Key(t *testing.T) {
	var f *File
	f = newFile(nil, "")
	parser1.parse(f)
	defaultSec := f.Section("")
	val := defaultSec.Key("key2")
	r := val.Val()
	expected := "1234adf"

	if expected != r {
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}

func TestFile_NotFound(t *testing.T) {
	var f *File
	f = newFile(nil, "")
	parser1.parse(f)
	_, err := f.section("adfdfadf")

	if err == nil {
		t.Errorf("expected error %s", err.Error())
	}
}
