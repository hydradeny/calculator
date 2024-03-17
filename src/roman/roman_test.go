package roman

import "testing"

func TestRomanToArabValid(t *testing.T) {
	testcases := []struct {
		in   string
		want int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"MMMCMXCIX", 3999},
		{"IV", 4},
		{"V", 5},
		{"IX", 9},
		{"X", 10},
		{"XIV", 14},
		{"XL", 40},
		{"L", 50},
		{"XC", 90},
		{"C", 100},
		{"CD", 400},
		{"D", 500},
		{"CM", 900},
		{"M", 1000},
		{"CDXLIV", 444},
		{"DCXXXIII", 633},
	}

	for _, tc := range testcases {
		res, err := RomanToArab(tc.in)
		if err != nil {
			t.Errorf("Error occurred: %s", err)
		}
		if res != tc.want {
			t.Errorf("RomanToArab: %v, want %v", res, tc.want)
		}
	}
}

func TestRomanToArabInvalid(t *testing.T) {
	testcases := []string{
		"",
		"IIII",
		"i",
		"IIVM",
		" ",
		"XM",
		"VL",
		"LC",
		"VC",
		"XM",
		"LD",
		"ksdfk",
		"XXX+",
		"734",
		"Zk",
		"X4X",
		"3CC",
		"MMMMMMMM",
	}
	for _, tc := range testcases {
		res, err := RomanToArab(tc)
		if err == nil {
			t.Errorf("error is nil on input %q:, res = %v", tc, res)
		}
	}
}
