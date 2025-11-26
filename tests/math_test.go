package tests

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_"math"
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

func TestRound(t *testing.T) {
	result := math_pkg.Round(3.6)
	assert.Equal(t, float64(4), result, fmt.Sprintf("expected 4 but got %f", result))
	assert.NotEqual(t, 16, result, fmt.Sprintf("did not expect 16 but got %f", result))
	fmt.Println("TestRound passed")
}

func TestIsNan(t *testing.T) {
	result := math_pkg.IsNan(4)
	require.Equal(t, true, result, "expected true but got false")
	fmt.Println("TestIsNan passed")
}