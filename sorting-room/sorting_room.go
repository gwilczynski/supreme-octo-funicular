package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %0.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %0.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	parsedNumber, _ := strconv.Atoi(fnb.Value())

	return parsedNumber
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	var number float64
	str, ok := fnb.(FancyNumber)

	if ok {
		number = float64(ExtractFancyNumber(str))
	} else {
		number = 0.0
	}

	return fmt.Sprintf("This is a fancy box containing the number %0.1f", number)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch value := i.(type) {
	case int:
		return DescribeNumber(float64(value))
	case float64:
		return DescribeNumber(value)
	case FancyNumber:
		return DescribeFancyNumberBox(value)
	case NumberBox:
		return DescribeNumberBox(value)
	case string:
		return "Return to sender"
	default:
		return DescribeFancyNumberBox(FancyNumber{""})
	}
}
