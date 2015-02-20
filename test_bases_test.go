package test_bases

import "testing"

func BenchmarkSearch1000(b *testing.B) {
	s := make([]int, 1000)
	for n := 0; n < b.N; n++ {
		Search(s, 3)
	}
}
