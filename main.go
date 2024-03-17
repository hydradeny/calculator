package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hydradeny/test_task/calculator/src/roman"
)

const (
	arabicNumChars = "1234567890"
	romanNumChars  = "IXVDLCM"
	operChars      = "+-*/"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите число, оператор(+-*/), число, разделенные пробелами:")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = text[0 : len(text)-1]
	result := Calculate(text)
	fmt.Println("Output:")
	fmt.Println(result)
}

func Calculate(text string) string {
	parts, err := parse(text)
	if err != nil {
		panic(err)
	}
	switch {
	// for arabic input
	case strings.Contains(arabicNumChars, parts[0][0:1]):
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		if a < 1 || a > 10 {
			panic("number(s) out of range")
		}
		b, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		if b < 1 || b > 10 {
			panic("number(s) out of range")
		}
		c, err := calculate(a, b, parts[1])
		if err != nil {
			panic(err)
		}
		return strconv.Itoa(c)
		// for roman input
	case strings.Contains(romanNumChars, parts[0][0:1]):
		a, err := roman.RomanToArab(parts[0])
		if err != nil {
			panic(err)
		}
		if a < 1 || a > 10 {
			panic("number(s) out of range")
		}
		b, err := roman.RomanToArab(parts[2])
		if err != nil {
			panic(err)
		}
		if b < 1 || b > 10 {
			panic("number(s) out of range")
		}
		c, err := calculate(a, b, parts[1])
		if err != nil {
			panic(err)
		}
		if c <= 0 {
			panic("Roman result cant be <=0")
		}
		result, err := roman.ArabToRoman(c)
		if err != nil {
			panic(err)
		}
		return result

	default:
		panic("Wrong format")
	}
}

func parse(text string) (result []string, err error) {
	result = strings.Split(text, " ")
	if len(result) != 3 {
		return result, fmt.Errorf("bad format")
	}
	return result, nil
}

func calculate(a int, b int, op string) (result int, err error) {
	switch op {
	case "-":
		return a - b, nil
	case "+":
		return a + b, nil
	case "/":
		return a / b, nil
	case "*":
		return a * b, nil
	}
	return result, fmt.Errorf("unknown operation")
}
