package main

import "testing"

func TestWcBytesCount(t *testing.T) {
	t.Parallel()
	want := 342190
	got := BytesCount()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
