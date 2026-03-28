package strutil

import (
	"strings"
	"testing"
)

func TestB64Encode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "aGVsbG8="},
		{"", ""},
		{"Hello, World!", "SGVsbG8sIFdvcmxkIQ=="},
		{"中文", "5Lit5paH"},
	}
	for _, tc := range tests {
		got := B64Encode(tc.input)
		if got != tc.want {
			t.Errorf("B64Encode(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestB64Decode(t *testing.T) {
	tests := []struct {
		input   string
		want    string
		wantErr bool
	}{
		{"aGVsbG8=", "hello", false},
		{"", "", false},
		{"SGVsbG8sIFdvcmxkIQ==", "Hello, World!", false},
		{"5Lit5paH", "中文", false},
		{"!!!invalid!!!", "", true},
	}
	for _, tc := range tests {
		got, err := B64Decode(tc.input)
		if (err != nil) != tc.wantErr {
			t.Errorf("B64Decode(%q) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			continue
		}
		if !tc.wantErr && got != tc.want {
			t.Errorf("B64Decode(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

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
