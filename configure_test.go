package configure

import (
	"fmt"
	"log"
	"testing"
)

func defaultListen(filepath string) {
	log.Println(fmt.Sprintf("file '%s' has been changed", filepath))
}

func TestWatch(t *testing.T) {
	file, _ := Watch("example1.txt", defaultListen)
	r1 := file.Filename()
	e1 := "example1.txt"
	// 输出：http
	r2 := file.Section("server").Key("protocol").Val()
	e2 := "http"

	r3 := file.Section("").Key("app_mode").Description()
	e3 := "possible values : production, development"

	if r1 != e1 || r2 != e2 || r3 != e3 {
		t.Errorf("expected '%s', '%s', '%s' but got '%s', '%s', '%s'", e1, e2, e3, r1, r2, r3)
	}
}
func TestWatchWithOption(t *testing.T) {
	file, _ := WatchWithOption("example2.txt", defaultListen, Option{Separation: "\t"})
	r1 := file.Filename()
	e1 := "example2.txt"
	// 输出：http
	r2 := file.Section("server").Key("protocol").Val()
	e2 := "http"

	r3 := file.Section("").Key("app_mode").Description()
	e3 := "possible values : production, development"

	if r1 != e1 || r2 != e2 || r3 != e3 {
		t.Errorf("expected '%s', '%s', '%s' but got '%s', '%s', '%s'", e1, e2, e3, r1, r2, r3)
	}
}
