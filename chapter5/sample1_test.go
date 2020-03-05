package chapter5

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	result, err := Division(4, 2)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 2 {
		t.Fatal("failed test")
	}
}

func TestSum(t *testing.T) {
	t.Run("引数1つ", func(t *testing.T) {
		t.Parallel()
		if Sum(1) != 1 {
			t.Fail()
		}
	})

	t.Run("引数2つ", func(t *testing.T) {
		t.Parallel()
		time.Sleep(time.Second * 1)
		if Sum(1, 2) != 3 {
			t.Fail()
		}
	})

	t.Run("引数3つ", func(t *testing.T) {
		t.Parallel()
		if Sum(1, 2, 3) != 6 {
			t.Fail()
		}
	})
}

func TestClock2_AddHour(t *testing.T) {
	now := time.Date(2019, 7, 04, 18, 30, 00, 0, time.Local)
	clock := Clock2{
		Now: func() time.Time { return now },
	}
	assert.Equal(t, now.Add(time.Hour*3), clock.AddHour(3))
}
