package main

import (
	"fmt"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
	// 実装案 1
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil

	//実装案 3
	//if !utf8.ValidString(s) {
	//	return s, errors.New("input is not valid UTF-8")
	//}

	// 実装案 2
	////fmt.Printf("input: %q\n", s)
	//r := []rune(s)
	////fmt.Printf("runes: %q\n", r)
	//for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
	//	r[i], r[j] = r[j], r[i]
	//}
	//return string(r), nil
}
