package main

import "testing"

type Test struct {
	word     string
	expected string
}

func DeleteVowels(s string) string

func TestDeleteVowels(t *testing.T) {
	cases := []Test{
		{"Golang", "Glng"},
		{"testing", "tstng"},
		{"DeleteVowels", "DltVwls"},
	}
	for _, cas := range cases {
		result := DeleteVowels(cas.word)
		if result != cas.expected {
			t.Fatalf("expected %v; got %v", cas.expected, result)
		}
	}
}
