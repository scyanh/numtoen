package main

import (
	"github.com/scyanh/numtoen/app/application/numberService"
	"testing"
)

func TestTranslateNumber(t *testing.T) {
	t.Logf("Process... TestTranslateNumber")

	var mapNumbers = map[string]string{
		"12345678": "twelve million three hundred forty five thousand six hundred seventy eight",
		"9223372036854775807": "nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven",
		"-9223372036854775808": "minus nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred eight",
	}
	for number, want := range mapNumbers {
		translated, _ := numberService.TranslateNumber(number)
		if translated != want {
			t.Errorf("want... \n %s \n but got... \n %s", want, translated)
		}
	}
}
