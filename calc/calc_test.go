package calc

import "testing"

type Test struct {
	expression string
	expected   float64
	wantError  bool
}

func TestCalc(t *testing.T) {
	tests := []Test{
		{"5+3", 8, false},
		{"8/0", 0, true},
		{"35 + (10 - 2 * 5) + (6 / 2 * 5 - 10 + 2) * (2 * 3)", 77, false},
		{"10 + 15 - (2 + 3) * 2", 15, false},
		{"5 * 8 + 4 * 6 + 15 - 14", 65, false},
		{"5-2+6*2-4+8/2-1", 14, false},
		{"35 + (10 - 2 * 5) + (6 / 0 * 5 - 10 + 2) * (2 * 3)", 0, true},
		{"(11437 + 128 * 31) / 237 - 37", 28, false},
		{"93478+23657-(52253/0)", 0, true},
		{"(37 296 / 37 - 17780 / 35) / 250", (37296/37 - 17780/35) / 250, false},
		{"5+(5-4", 0, true},
		{"35 + (10 - 2 * 5) + (6 / 3 * 5 - 10 + 2 * (2 * 3)", 0, true},
		{"5*+7", 0, true},
		{"5+7/", 0, true},
		{"valid input", 0, true},
		{"", 0, true},
	}

	for i, test := range tests {
		result, err := Calc(test.expression)
		if test.wantError && err == nil {
			t.Fatalf("Expected error; got %v in test number %v", result, i+1)
		}
		if !test.wantError && err != nil {
			t.Fatalf("Unexpected error in test #%v: %v", i+1, err)
		}
		if result != test.expected {
			t.Fatalf("Expected %v; got %v in test # %v", test.expected, result, i+1)
		}
	}
}
