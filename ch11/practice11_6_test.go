package main

import (
	"testing"

	"gobook/ch2/popcount"
	"gobook/ch2/practice2_4"
	"gobook/ch2/practice2_5"
)

func BenchmarkProcCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(100)
		popcount.PopCount(10)
		popcount.PopCount(1000)
	}
}

func BenchmarkPractice2_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice2_4.PopCount(100)
		practice2_4.PopCount(10)
		practice2_4.PopCount(1000)
	}
}

func BenchmarkPractice2_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice2_5.PopCount(100)
		practice2_5.PopCount(10)
		practice2_5.PopCount(1000)
	}
}
