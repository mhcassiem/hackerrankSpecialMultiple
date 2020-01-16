package tests

import (
	"github.com/mhcassiem/SpecialMultiple/pkg"
	"testing"
)

func TestCalculator(t *testing.T)  {

	result := pkg.Calculator

	if result(1) != 1 {
		t.Errorf("Expected %d, received %d", 1, result(1))
	}

}