package chapter5

import (
	"fmt"
	"time"
)

// Division 割り算
func Division(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("invalid num=%v", y)
	}
	return x / y, nil
}

func Sum(n ...int) int {
	var sum int
	for _, m := range n {
		sum += m
	}

	return sum
}

// AddHour 今の時間にｘ時間を加算
func AddHour(hour int) time.Time {
	return time.Now().Add(time.Hour * time.Duration(hour))
}

// Before -------------------------
type Clock1 struct{}

// AddHour 今の時間にｘ時間を加算
func (c Clock1) AddHour(hour int) time.Time {
	return time.Now().Add(time.Hour * time.Duration(hour))
}

// After -------------------------
type Clock2 struct {
	Now func() time.Time
}

// AddHour 今の時間にｘ時間を加算
func (c Clock2) AddHour(hour int) time.Time {
	return c.Now().Add(time.Hour * time.Duration(hour))
}
