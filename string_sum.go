package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var sum int64
	ps := strings.ReplaceAll(input, " ", "")
	if ps == "" {
		return "", fmt.Errorf("got empty string: %w", errorEmptyInput)
	}
	ps = strings.ReplaceAll(ps, "-", "+-")
	strSlice := strings.FieldsFunc(ps, func(r rune) bool {
		return r == '+'
	})
	if len(strSlice) != 2 {
		return "", fmt.Errorf("expecting exactly 2 operands: %w", errorNotTwoOperands)
	}
	for _, si := range strSlice {
		ii, err := strconv.ParseInt(si, 10, 64)
		if err != nil {
			return "", fmt.Errorf("error while processing %s: %w", si, err)
		}
		sum += ii
	}

	return strconv.FormatInt(sum, 10), nil
}
