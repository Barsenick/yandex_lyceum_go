package getutflength

import (
	"testing"
)

type Test struct {
	input     []byte
	expected  int
	wantError bool
}

var invalidUTF8 = []byte{
	0xC0, 0x80,
}

func TestGetUTFLength(t *testing.T) {
	cases := []Test{
		{[]byte("Golang"), 6, false},
		{[]byte("testing"), 7, false},
		{invalidUTF8, 0, true},
	}
	for _, test := range cases {
		result, err := GetUTFLength(test.input)
		if result != test.expected {
			t.Fatalf("expected %v; got %v", test.expected, result)
		}
		if err == nil && test.wantError {
			t.Fatalf("expected error")
		}
		if err != nil && !test.wantError {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}
