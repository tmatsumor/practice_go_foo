package foo

import "testing"

func TestFoo(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
