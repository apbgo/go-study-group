package chapter1

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	t.Run("Calc", func(t *testing.T) {
		a, err := Calc("+", 1, 5)
		assert.NoError(t, err)
		assert.Equal(t, 6, a)

		b, err := Calc("-", 4, 2)
		assert.NoError(t, err)
		assert.Equal(t, 2, b)

		c, err := Calc("×", 2, 5)
		assert.NoError(t, err)
		assert.Equal(t, 10, c)

		d, err := Calc("÷", 5, 2)
		assert.NoError(t, err)
		assert.Equal(t, 2, d)

		e, err := Calc("@", 5, 2)
		assert.Error(t, err)
		assert.Equal(t, 0, e)
	})
}

func TestStringEncode(t *testing.T) {
	t.Run("StringEncode", func(t *testing.T) {
		assert.Equal(t, "SnAk", StringEncode("sn_ak"))
		assert.Equal(t, "as_h_gkd_jahf", StringEncode("asHGkdJahf"))
	})
}

func TestSqrt(t *testing.T) {
	t.Run("Sqrt", func(t *testing.T) {
		assert.Equal(t, 3.0, Sqrt(9.0))
		assert.Equal(t, true, 0.00000000000005 > math.Abs(math.Sqrt(38.0)-Sqrt(38.0)))
	})
}

func TestPyramid(t *testing.T) {
	t.Run("Pyramid", func(t *testing.T) {
		assert.Equal(t, "1\n12\n123\n1234\n12345", Pyramid(5))
		assert.Equal(t, "1\n12\n123\n1234\n12345\n123456\n1234567\n12345678\n123456789\n12345678910", Pyramid(10))
	})
}

func TestStringSum(t *testing.T) {
	t.Run("StringSum", func(t *testing.T) {
		i, err := StringSum("2", "10")
		assert.NoError(t, err)
		assert.Equal(t, 12, i)
		j, err := StringSum("aaa", "5")
		assert.Error(t, err)
		assert.Equal(t, 0, j)
	})
}

func TestSumFromFileNumber(t *testing.T) {
	t.Run("SumFromFileNumber", func(t *testing.T) {
		i, err := SumFromFileNumber("test/numbers.txt")
		assert.NoError(t, err)
		assert.Equal(t, 55, i)
	})
}
