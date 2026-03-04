package strutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

const (
	alphaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numChars   = "0123456789"
	mixChars   = alphaChars + numChars
)

// Hash computes a hash of text using the given algorithm (md5, sha256, sha1).
func Hash(algo, text string) (string, error) {
	switch strings.ToLower(algo) {
	case "md5":
		h := md5.Sum([]byte(text))
		return hex.EncodeToString(h[:]), nil
	case "sha256":
		h := sha256.Sum256([]byte(text))
		return hex.EncodeToString(h[:]), nil
	case "sha1":
		h := sha1.Sum([]byte(text))
		return hex.EncodeToString(h[:]), nil
	default:
		return "", fmt.Errorf("unsupported hash type: %s (use md5, sha256, sha1)", algo)
	}
}

// B64Encode encodes text to base64.
func B64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// B64Decode decodes base64 text.
func B64Decode(text string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed: %w", err)
	}
	return string(b), nil
}

// URLEncode percent-encodes text.
func URLEncode(text string) string {
	return url.QueryEscape(text)
}

// URLDecode decodes a percent-encoded string.
func URLDecode(text string) (string, error) {
	s, err := url.QueryUnescape(text)
	if err != nil {
		return "", fmt.Errorf("url decode failed: %w", err)
	}
	return s, nil
}

// UUID generates a random UUID v4.
func UUID() string {
	return uuid.New().String()
}

// RandString generates a random string of length n with the given character set type.
// typ: "alpha", "num", "mix"
func RandString(n int, typ string) (string, error) {
	var charset string
	switch strings.ToLower(typ) {
	case "alpha":
		charset = alphaChars
	case "num":
		charset = numChars
	case "mix":
		charset = mixChars
	default:
		return "", fmt.Errorf("unsupported type: %s (use alpha, num, mix)", typ)
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b), nil
}

// ToSnake converts text to snake_case.
func ToSnake(text string) string {
	var result []rune
	runes := []rune(text)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(runes[i-1]) || (i+1 < len(runes) && unicode.IsLower(runes[i+1]))) {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else if r == '-' || r == ' ' {
			result = append(result, '_')
		} else {
			result = append(result, r)
		}
	}
	s := string(result)
	// collapse multiple underscores
	for strings.Contains(s, "__") {
		s = strings.ReplaceAll(s, "__", "_")
	}
	return strings.Trim(s, "_")
}

// ToCamel converts text to camelCase.
func ToCamel(text string) string {
	s := toPascal(text)
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// ToPascal converts text to PascalCase.
func ToPascal(text string) string {
	return toPascal(text)
}

func toPascal(text string) string {
	words := splitWords(text)
	var result strings.Builder
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		runes := []rune(w)
		result.WriteRune(unicode.ToUpper(runes[0]))
		for _, r := range runes[1:] {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

// splitWords splits a string by underscores, hyphens, spaces, and camelCase boundaries.
func splitWords(text string) []string {
	var words []string
	var current []rune
	runes := []rune(text)
	for i, r := range runes {
		if r == '_' || r == '-' || r == ' ' {
			if len(current) > 0 {
				words = append(words, string(current))
				current = current[:0]
			}
			continue
		}
		if i > 0 && unicode.IsUpper(r) && unicode.IsLower(runes[i-1]) {
			words = append(words, string(current))
			current = []rune{r}
			continue
		}
		current = append(current, r)
	}
	if len(current) > 0 {
		words = append(words, string(current))
	}
	return words
}
