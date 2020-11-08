package numberService

import (
	"errors"
	"math"
	"strconv"
)

// how many digit's groups to process
const groupsNumber int = 7

var smallNumbers = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}
var tenNumbers = []string{
	"", "", "twenty", "thirty", "forty", "fifty",
	"sixty", "seventy", "eighty", "ninety",
}
var scaleNumbers = []string{
	"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
}

func TranslateNumber(numberString string) (string, error) {
	number, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		return "", errors.New("wrong number, please provide a number between -9223372036854775808 and 9223372036854775807")
	}

	return convert(number, numberString), nil
}

func convert(number int64, numberString string) string {
	// Zero rule
	if number == 0 {
		return smallNumbers[0]
	}

	// If negative number string, remove minus
	elementString := string(numberString[0])
	if elementString == "-" {
		numberString = numberString[1:]
	}

	// Divide into three-digits group
	var groups [groupsNumber]int

	j := 0
	groupIdx := 0
	numberByGroup := 0
	for i := range numberString {
		numberFromString, _ := strconv.Atoi(string(numberString[len(numberString)-i-1]))
		numberByGroup += numberFromString * int(math.Pow10(j))

		j++
		if j == 3 || i == len(numberString)-1 {
			groups[groupIdx] = numberByGroup
			groupIdx++
			j = 0
			numberByGroup = 0
		}
	}

	// Translate to text
	var textGroup [groupsNumber]string
	for i := 0; i < groupsNumber; i++ {
		textGroup[i] = digitGroup2Text(groups[i])
	}
	combined := textGroup[0]

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			prefix := textGroup[i] + " " + scaleNumbers[i]

			if len(combined) != 0 {
				prefix += " "
			}

			combined = prefix + combined
		}
	}

	// attach "minus" if it is negative number
	if number < 0 {
		combined = "minus " + combined
	}

	return combined
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func digitGroup2Text(group int) (ret string) {
	hundreds := group / 100
	tensUnits := intMod(group, 100)

	if hundreds != 0 {
		ret += smallNumbers[hundreds] + " hundred"

		if tensUnits != 0 {
			ret += " "
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		ret += tenNumbers[tens]

		if units != 0 {
			ret += " " + smallNumbers[units]
		}
	} else if tensUnits != 0 {
		ret += smallNumbers[tensUnits]
	}

	return
}
