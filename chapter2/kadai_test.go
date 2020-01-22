package chapter2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	t.Run("Calc", func(t *testing.T) {
		i0, err := Calc([]int{})
		assert.Error(t, err)
		assert.Equal(t, 0, i0)

		i1, err := Calc([]int{3})
		assert.NoError(t, err)
		assert.Equal(t, 3, i1)

		i2, err := Calc([]int{3, 7})
		assert.NoError(t, err)
		assert.Equal(t, 21, i2)

		i3, err := Calc([]int{3, 7, 10})
		assert.NoError(t, err)
		assert.Equal(t, 20, i3)

		i4, err := Calc([]int{3, 7, 10, 21})
		assert.NoError(t, err)
		assert.Equal(t, 41, i4)
	})
}

func TestNumbers(t *testing.T) {
	t.Run("Numbers", func(t *testing.T) {
		i := Numbers()
		assert.Equal(t, []Number{
			{
				index: 1,
			},
			{
				index: 2,
			},
			{
				index: 3,
			},
		}, i)
	})
}

func TestCalcMap(t *testing.T) {
	t.Run("CalcMap", func(t *testing.T) {
		i1 := CalcMap(map[string]int{
			"ichi": 1,
			"ni":   2,
			"san":  3,
			"yon":  4,
			"go":   5,
		})

		assert.Equal(t, 11, i1)

		i2 := CalcMap(map[string]int{
			"roku":    6,
			"nana":    7,
			"zyuuyon": 14,
			"go":      5,
		})

		assert.Equal(t, 32, i2)
	})
}

func TestAdd(t *testing.T) {
	var models = []Model{
		{Value: 1},
		{Value: 5},
		{Value: 11},
	}

	t.Run("Add", func(t *testing.T) {
		Add(models)

		assert.Equal(t, []Model{
			{Value: 6},
			{Value: 10},
			{Value: 16},
		}, models)
	})
}

func TestUnique(t *testing.T) {
	t.Run("Unique", func(t *testing.T) {
		u1 := Unique([]int{11, 11, 11, 2, 2})
		assert.Equal(t, []int{11, 2}, u1)

		u2 := Unique([]int{1, 3, 3, 3, 3, 3})
		assert.Equal(t, []int{1, 3}, u2)
	})
}

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	t.Run("Fibonacci", func(t *testing.T) {
		f := Fibonacci()
		for i := 0; i < 10; i++ {
			assert.Equal(t, expected[i], f())
		}

	})
}
