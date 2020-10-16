package configure

import "testing"

var contentWithDefault = []string{"# test1","key = 1", "# test2", "key2 = 2"}
//var contentWithoutDefault = []string{"[sec0]","# test1","key = 1","[sec1]","key2 = aaa","key3 = bbb","key4 = ccc","[sec2]","# desdes","key5=adsfeq1234"}
var contentEmpty = []string{}


func TestSection_Name(t *testing.T) {
	testSection,_ := NewSection("test", nil, nil)
	r := testSection.Name()
	expected := "test"
	if r != expected {
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}

func TestSection_ItemKeyVal(t *testing.T) {
	testSection,_ := NewSection("test", contentWithDefault, nil)
	len_r := len(testSection.items)

	firstName := testSection.items[0].Name()

	expected_r := 2
	expected_name := "key1"

	if expected_r != len_r && expected_name != firstName {
		t.Errorf("expected length '%d','%s' but got, '%d' '%s'", expected_r, expected_name, len_r, firstName)
	}
}

func TestSection_Description(t *testing.T) {
	var testSection,_ = NewSection("test", contentWithDefault, nil)
	len_r := len(testSection.items)

	firstName := testSection.items[1].Description()

	expected_r := 2
	expected_name := "test2"

	if expected_r != len_r && expected_name != firstName {
		t.Errorf("expected length '%d','%s' but got, '%d' '%s'", expected_r, expected_name, len_r, firstName)
	}
}
