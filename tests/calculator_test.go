package tests

import (
	"github.com/mhcassiem/hackerrankSpecialMultiple/pkg"
	"testing"
)

func TestGetMultiple(t *testing.T) {

	var expected = "9"
	var input int32 = 1
	if expected != runTestCase(input) {
		t.Errorf("Expected %s, received %s", expected, runTestCase(input))
	}

	expected = "9"
	input = 3
	if expected != runTestCase(input) {
		t.Errorf("Expected %s, received %s", expected, runTestCase(input))
	}

	expected = "90"
	input = 5
	if expected != runTestCase(input) {
		t.Errorf("Expected %s, received %s", expected, runTestCase(input))
	}

	expected = "9009"
	input = 7
	if expected != runTestCase(input){
		t.Errorf("Expected %s, received %s", expected, runTestCase(input))
	}

}

func runTestCase(input int32) string {
	result := pkg.GetMultiple(input)
	return result
}
