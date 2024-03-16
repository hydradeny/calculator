package roman

import (
	"fmt"
)

var romanar = [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
var arabar = [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

// Function converts a string which contains roman number,
// formed by classical rules. So number can be in range
// from 1 to 3999 included
func RomanToArab(roman string) (arab int, err error) {
	if roman == "" {
		return arab, fmt.Errorf("empty string")
	}
	c := 0
	i := 0
	for n := len(romanar) - 1; n >= 0 && i < len(roman); {
		pir := n%2 == 1
		if roman[i] == romanar[n][0] && (!pir || (i+1 != len(roman) && roman[i+1] == romanar[n][1])) {
			arab += arabar[n]
			if pir {
				n--
				i += 2
			} else {
				i++
				c++
				if c >= 3 {
					n--
					c = 0
				}
			}
		} else {
			n--
			c = 0
		}
	}
	if i < len(roman) {
		return arab, fmt.Errorf("bad roman number")
	}
	return arab, nil
}

// Function converts arabic number to roman number according to
// classical rules. So value can be in range from 1 to 3999 included
func ArabToRoman(arab int) (roman string, err error) {
	if arab <= 0 || arab > 3999 {
		return roman, fmt.Errorf("bad arabic number")
	}
	n := len(arabar) - 1
	var result []byte
	for arab > 0 {
		if arab >= arabar[n] {
			result = append(result, romanar[n]...)
			arab -= arabar[n]

		} else {
			n--
		}
	}
	return string(result), nil
}
