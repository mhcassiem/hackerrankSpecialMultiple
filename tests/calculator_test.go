package tests

import (
	"github.com/mhcassiem/hackerrankSpecialMultiple/pkg"
	"testing"
)

func TestGetMultiple(t *testing.T) {

	var expected = "9"
	var input int32 = 1
	if expected != returns9for1(input) {
		t.Errorf("Expected %s, received %s", expected, returns9for1(input))
	}

	expected = "9"
	input = 3
	if expected != returns9for3(input) {
		t.Errorf("Expected %s, received %s", expected, returns9for1(input))
	}

	expected = "90"
	input = 5
	if expected != returns90for5(input) {
		t.Errorf("Expected %s, received %s", expected, returns9for1(input))
	}

	expected = "9009"
	input = 7
	if expected != returns9009for7(input){
		t.Errorf("Expected %s, received %s", expected, returns9for1(input))
	}

}

func returns9for1(input int32) string {
	result := pkg.GetMultiple(input)
	return result
}

func returns9for3(input int32) string {
	result := pkg.GetMultiple(input)
	return result
}

func returns90for5(input int32) string {
	result := pkg.GetMultiple(input)
	return result
}

func returns9009for7(input int32) string {
	result := pkg.GetMultiple(input)
	return result
}
