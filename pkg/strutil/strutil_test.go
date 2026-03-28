package strutil

import (
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
