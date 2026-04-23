package greeting_test

import (
	"testing"

	"github.com/pboyd/hello/internal/greeting"
)

func TestMessage(t *testing.T) {
	got := greeting.Message()
	want := "Hello, World!"
	if got != want {
		t.Errorf("Message() = %q, want %q", got, want)
	}
}
