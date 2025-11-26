package tests

import (
	"fmt"
	_ "math"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Skipping test on macOS")
	}
	result := math_pkg.Power(3, 3)
	require.Equal(t, float64(27), result, "expected 27 but got different value")
	fmt.Println("TestSkip passed")
}

func TestMain(m *testing.M) {
	fmt.Println("Starting tests...")
	m.Run()
	fmt.Println("Tests completed.")
}

func TestSubTest(t *testing.T) {
	t.Run("Should be test math is nan", func(t *testing.T) {
		result := math_pkg.IsNan(0.4)
		require.Equal(t, false, result, "expected false but got true")
		fmt.Println("SubTest IsNan passed")
	})

	t.Run("Should be test math power", func(t *testing.T) {
		result := math_pkg.Power(10, 2)
		require.Equal(t, float64(100), result, fmt.Sprintf("expected 100 but actual value is %v", result))
		fmt.Println("SubTest Power passed")
	})
}

func TestMathTable(t *testing.T) {
	// declare tests array of struct
	tests := []struct {
		name string
		request float64
		expected float64
		option string
	} {
		{
			name: "Test Square Root of 49",
			request: 49,
			expected: 7,
			option: "sqrt",
		},
		{
			name: "Test Square Root of 64",
			request: 64,
			expected: 8,
			option: "sqrt",
		},
		{
			name: "Test Round of 4.4",
			request: 4.4,
			expected: 4,
			option: "round",
		},
	}

	// iterate through tests array of struct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// declare result
			var result float64
			// option check
			switch test.option {
			case "sqrt":
				result = math_pkg.SquareRoot(test.request)
			case "round":
				result = math_pkg.Round(test.request)
			}
			// assertion
			require.Equal(t, test.expected, result, fmt.Sprintf("expected %f but got %f", test.expected, result))
			// print if success
			fmt.Println("Table Test passed for:", test.name)
		})
	}
}