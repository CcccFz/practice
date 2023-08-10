package main

import "testing"

func BenchmarkWriteFileDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteFile("a.txt", false)
	}
}

func BenchmarkWriteFileBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteFile("b.txt", true)
	}
}

func BenchmarkWriteFileOnetime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneWriteFile("c.txt", genBigBytes())
	}
}
