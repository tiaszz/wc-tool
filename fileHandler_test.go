package main

import "testing"

const test = "test.txt"

func TestWcBytesCount(t *testing.T) {
	t.Parallel()

	file := OpenFile(test)
	defer file.Close()

	want := 342190
	got := BytesCount(file)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWcLinesCount(t *testing.T) {
	t.Parallel()

	file := OpenFile(test)
	defer file.Close()

	want := 7145
	got := LinesCount(file)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWcCountWords(t *testing.T) {
	t.Parallel()

	file := OpenFile(test)
	defer file.Close()

	want := 58164
	got := CountWords(file)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
