package main

import (
	"os"
	"testing"
)

var testdata_path = "../../testdata/"

func Test_CountPixels(t *testing.T) {

	// Arrange
	testfile1 := open(testdata_path + "test file 1.raw")
	t.Cleanup(func() { testfile1.Close() })
	want := 10

	// Act
	got := int(pixelCount(testfile1))

	// Assert
	if got != want {
		t.Errorf("%s failed : expected [%v] but got [%v]", t.Name(), want, got)
	}
}

func Test_Filelist(t *testing.T) {

	// Arrange
	dirlist, err := os.ReadDir(testdata_path)
	if err != nil {
		t.Errorf("%s unable to read testdata from [%s] : %v", t.Name(), testdata_path, err)
	}

	t.Run("Test_MakeFilelist_Length", func(t *testing.T) {

		// Arrange
		want := len(dirlist)
		directory := open(testdata_path)
		t.Cleanup(func() { directory.Close() })

		// Act
		got := makeFilelist(directory)

		// Assert
		if len(got) != want {
			t.Errorf("%s failed : expected [%v] but got [%v]", t.Name(), want, got)
		}
	})

	t.Run("Test_FilterFilelist_Length", func(t *testing.T) {

		// Arrange
		want := len(dirlist) - 1
		directory := open(testdata_path)
		t.Cleanup(func() { directory.Close() })
		list := makeFilelist(directory)
		exclude := open(testdata_path + list[0].Name())

		// Act
		got := len(filterFilelist(list, exclude))

		// Assert
		if got != want {
			t.Errorf("%s failed : expected [%v] but got [%v]", t.Name(), want, got)
		}
	})

	t.Run("Test_FilterFilelist_Contents", func(t *testing.T) {

		// Arrange
		want := dirlist[1].Name()
		directory := open(testdata_path)
		t.Cleanup(func() { directory.Close() })
		list := makeFilelist(directory)
		exclude := open(testdata_path + list[0].Name())

		// Act
		got := filterFilelist(list, exclude)[0].Name()

		// Assert
		if got != want {
			t.Errorf("%s failed : expected [%v] but got [%v]", t.Name(), want, got)
		}
	})
}

func BenchmarkTest(b *testing.B) {

	for i := 0; i < b.N; i = i + 1 {

	}
}
