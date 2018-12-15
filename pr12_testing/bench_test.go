package main

import "testing"

// Benchmark test run by go test -v -run x -bench .
func BenchmarkDecoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decoder("jsonfile.json")
	}
}

func BenchmarkUnmarshaling(b *testing.B) {
	for i:=0; i<b.N; i++{
		unmarshaling("jsonfile.json`")
	}
}