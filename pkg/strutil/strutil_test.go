package strutil

import (
	"strings"
	"testing"
)

func TestUUID(t *testing.T) {
	id := UUID()
	if len(id) != 36 {
		t.Fatalf("expected length 36, got %d: %s", len(id), id)
	}
	if strings.Count(id, "-") != 4 {
		t.Fatalf("expected 4 dashes, got: %s", id)
	}
}

func TestUUIDV1(t *testing.T) {
	id, err := UUIDV1()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(id) != 36 {
		t.Fatalf("expected length 36, got %d: %s", len(id), id)
	}
	if strings.Count(id, "-") != 4 {
		t.Fatalf("expected 4 dashes, got: %s", id)
	}
}

func TestFormatUUID(t *testing.T) {
	id := "550e8400-e29b-41d4-a716-446655440000"

	tests := []struct {
		noDash bool
		upper  bool
		want   string
	}{
		{false, false, "550e8400-e29b-41d4-a716-446655440000"},
		{true, false, "550e8400e29b41d4a716446655440000"},
		{false, true, "550E8400-E29B-41D4-A716-446655440000"},
		{true, true, "550E8400E29B41D4A716446655440000"},
	}

	for _, tt := range tests {
		got := FormatUUID(id, tt.noDash, tt.upper)
		if got != tt.want {
			t.Errorf("FormatUUID(%v, %v) = %q, want %q", tt.noDash, tt.upper, got, tt.want)
		}
	}
}
