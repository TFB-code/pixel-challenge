package main

import "testing"

func TestCountPixels(t *testing.T) {

	// Arrange
	testfile1 := open("../../testdata/test file 1.raw")
	defer testfile1.Close()
	want := 10

	// Act
	got := int(pixelCount(testfile1))

	// Assert
	if got != want {
		t.Errorf("TestCountPixels failed : expected %v but got %v", want, got)
	}
}

func BenchmarkTest(b *testing.B) {

	for i := 0; i < b.N; i = i + 1 {

	}
}
