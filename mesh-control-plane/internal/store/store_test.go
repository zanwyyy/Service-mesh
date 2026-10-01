package store

import "testing"

func TestStateSetGet(t *testing.T) {
	state := NewState()
	state.Set("service-a", "healthy")

	got, ok := state.Get("service-a")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if got != "healthy" {
		t.Fatalf("unexpected value: %s", got)
	}
}
