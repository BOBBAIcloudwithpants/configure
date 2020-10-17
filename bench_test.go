package configure

import (
	"testing"
)

func BenchmarkWatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Watch("example1.txt", defaultListen)
	}
}

func BenchmarkWatchWithOption(b *testing.B) {
	WatchWithOption("example2.txt", defaultListen, Option{Separation: "\t"})
}

func BenchmarkFile_Section(b *testing.B) {
	f := newFile(nil, "")
	parser1.parse(f)
	for i := 0; i < b.N; i++ {
		f.Section("")
	}
}
func BenchmarkFile_Key(b *testing.B) {
	f := newFile(nil, "")
	parser1.parse(f)
	for i := 0; i < b.N; i++ {
		f.Section("").Key("key2")
	}
}

func BenchmarkFile_Val(b *testing.B) {
	f := newFile(nil, "")
	parser1.parse(f)
	for i := 0; i < b.N; i++ {
		f.Section("").Key("key2").Val()
	}
}
