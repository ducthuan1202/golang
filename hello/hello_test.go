package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestWorld(t *testing.T) {
    want := "Bidgear"
    if got := SayHi(); got != want {
        t.Errorf("SayHi() = %q, want %q", got, want)
    }
}