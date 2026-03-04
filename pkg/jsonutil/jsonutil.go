package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// readInput reads from a file path or stdin if path is empty.
func readInput(filePath string) ([]byte, error) {
	if filePath != "" {
		return os.ReadFile(filePath)
	}
	return io.ReadAll(os.Stdin)
}

// Format pretty-prints JSON from filePath (or stdin if empty).
func Format(filePath string) (string, error) {
	data, err := readInput(filePath)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}
	return buf.String(), nil
}

// Minify compacts JSON from filePath (or stdin if empty).
func Minify(filePath string) (string, error) {
	data, err := readInput(filePath)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}
	var buf bytes.Buffer
	if err := json.Compact(&buf, data); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}
	return buf.String(), nil
}

// Validate checks whether the input is valid JSON.
// Returns nil if valid, error with details otherwise.
func Validate(filePath string) error {
	data, err := readInput(filePath)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	if !json.Valid(data) {
		// Try to get a meaningful error message
		var v interface{}
		if err := json.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("invalid JSON: %w", err)
		}
	}
	return nil
}
