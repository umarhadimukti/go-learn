package tests

import (
	"testing"
	"fmt"
	math_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/math"
)

func TestSquareRoot(t *testing.T) {
	if sqrt := math_pkg.SquareRoot(25); sqrt != 4 {
		t.Error("TestSquareRoot failed:", fmt.Sprintf("expected 4 but got %f", sqrt))
	}
	fmt.Println("TestSquareRoot passed")
}

func TestPower(t *testing.T) {
	if power := math_pkg.Power(2, 4); power != 16 {
		t.Fatal("TestPower failed:", fmt.Sprintf("expected 16 but got %f", power))
	}
	fmt.Println("TestPower passed")
}