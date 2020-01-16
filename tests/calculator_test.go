package tests

import (
	"github.com/mhcassiem/SpecialMultiple/pkg"
	"testing"
)

func TestGetMultiple(t *testing.T)  {

	var expected = 9
	var input = 1
	if expected != returns9for1(input){
		t.Errorf("Expected %d, received %d", expected, returns9for1(input))
	}

	expected = 90
	input = 5
	if expected != returns90for5(input){
		t.Errorf("Expected %d, received %d", expected, returns9for1(input))
	}

}

func returns9for1(input int) int {
	result := pkg.GetMultiple(input)
	return result
}

func returns90for5(input int) int {
	result := pkg.GetMultiple(input)
	return result
}