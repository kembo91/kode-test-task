package utils_test

import (
	"testing"

	"github.com/kembo91/kode-test-task/server/internal/utils"
)

func TestIsValidAnagram(t *testing.T) {
	var tests = []struct {
		teststr string
		exp     bool
	}{
		{"1122334", false},
		{"anagram", true},
		{"anagram22", false},
		{"ana_gram", false},
		{"ana22_gram", false},
		{"*#@@31", false},
		{"", false},
		{"a", true},
		{"1", false},
		{"ddddd", true},
		{"hello anagram", false},
	}
	for _, tv := range tests {
		var got bool
		err := utils.IsValidAnagram(tv.teststr)
		if err == nil {
			got = true
		} else {
			got = false
		}
		if got != tv.exp {
			t.Errorf(`Expected %v for %v got %v`, tv.exp, tv.teststr, got)
		}
	}
}

func TestIsValidUsername(t *testing.T) {
	var tests = []struct {
		teststr string
		exp     bool
	}{
		{"1122334", true},
		{"anagram", true},
		{"anagram22", true},
		{"ana_gram", true},
		{"ana22_gram", true},
		{"*#@@31", false},
		{"", false},
		{"a", false},
		{"1", false},
		{"ddddd", true},
		{"hello anagram", false},
	}
	for _, tv := range tests {
		var got bool
		err := utils.IsValidUsername(tv.teststr)
		if err == nil {
			got = true
		} else {
			got = false
		}
		if got != tv.exp {
			t.Errorf(`Expected %v for %v got %v`, tv.exp, tv.teststr, got)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	var tests = []struct {
		teststr string
		exp     bool
	}{
		{"123456789", true},
		{"12345", false},
	}
	for _, tv := range tests {
		var got bool
		err := utils.IsValidPassword(tv.teststr)
		if err == nil {
			got = true
		} else {
			got = false
		}
		if got != tv.exp {
			t.Errorf(`Expected %v for %v got %v`, tv.exp, tv.teststr, got)
		}
	}
}
