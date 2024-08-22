package auth

import (
	"reflect"
	"testing"
)

func TestSimple(t *testing.T) {
	got := ""
	want := ""
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
