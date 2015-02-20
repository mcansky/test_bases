package test_bases

import "testing"

func BenchmarkSearch1000(b *testing.B) {
	s := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = i + 1
	}

	for n := 0; n < b.N; n++ {
		Search(s, 3)
	}
}
