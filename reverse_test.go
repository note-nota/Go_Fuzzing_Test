package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	fmt.Println(" >>> TestReverse")

	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, revErr := Reverse(tc.in)
		if revErr != nil {
			fmt.Printf("Reverse: %q, err %v\n", rev, revErr)
			return
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	fmt.Println(" >>> FuzzReverse")

	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	// Fuzz test 実装案　1
	f.Fuzz(func(t *testing.T, orig string) {
		rev, _ := Reverse(orig)
		doubleRev, _ := Reverse(rev)

		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})

	// Fuzz test 実装案　2
	//f.Fuzz(func(t *testing.T, orig string) {
	//	rev, revErr := Reverse(orig)
	//	if revErr != nil {
	//		return
	//	}
	//
	//	doubleRev, doubleRevErr := Reverse(rev)
	//	if doubleRevErr != nil {
	//		return
	//	}
	//
	//	t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
	//	if orig != doubleRev {
	//		t.Errorf("Before: %q, reverse: %q, after: %q", orig, rev, doubleRev)
	//	}
	//	if utf8.ValidString(orig) && !utf8.ValidString(rev) {
	//		t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
	//	}
	//})
}
