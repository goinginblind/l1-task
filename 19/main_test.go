package main_test

import (
	"strings"
	"testing"

	main "github.com/goinginblind/l1-task/19"
)

// unit tests
func TestReverseStringUTF8(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "English",
			s:    "Hello World",
			want: "dlroW olleH",
		},
		{
			name: "Russian",
			s:    "главрыба",
			want: "абырвалг",
		},
		{
			name: "Chinese",
			s:    "你好世界",
			want: "界世好你",
		},
		{
			name: "Emojis",
			s:    "🍕🐹",
			want: "🐹🍕",
		},
		{
			name: "Mixed",
			s:    "глав🐟, 🐹con",
			want: "noc🐹 ,🐟валг",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.ReverseStringUTF8(tt.s)
			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("ReverseStringUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStringRunes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "English",
			s:    "Hello World",
			want: "dlroW olleH",
		},
		{
			name: "Russian",
			s:    "главрыба",
			want: "абырвалг",
		},
		{
			name: "Chinese",
			s:    "你好世界",
			want: "界世好你",
		},
		{
			name: "Emojis",
			s:    "🍕🐹",
			want: "🐹🍕",
		},
		{
			name: "Mixed",
			s:    "глав🐟, 🐹con",
			want: "noc🐹 ,🐟валг",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.ReverseStringRunes(tt.s)
			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("ReverseStringUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

// benchmarks
var benchString = "Hello 世界! 🐍 Let's test some длинные строки! " + strings.Repeat("🐟🇷🇺🖥️⛏️", 1000)

func BenchmarkReverseUTF8(b *testing.B) {
	for b.Loop() {
		_ = main.ReverseStringUTF8(benchString)
	}
}

func BenchmarkReverseRunes(b *testing.B) {
	for b.Loop() {
		_ = main.ReverseStringRunes(benchString)
	}
}

// goos: linux
// goarch: amd64
// pkg: github.com/goinginblind/l1-task/19
// cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx
// BenchmarkReverseUTF8-8    	   12661	     94250 ns/op	   27265 B/op	       1 allocs/op
// BenchmarkReverseRunes-8   	    4302	    277435 ns/op	   55937 B/op	       2 allocs/op
// PASS
// ok  	github.com/goinginblind/l1-task/19	2.393s
