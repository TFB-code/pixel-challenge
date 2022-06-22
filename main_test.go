package main

import "testing"

func TestNothing(t *testing.T) {}

func BenchmarkTest(b *testing.B) {

	for i := 0; i < b.N; i = i + 1 {

	}
}
