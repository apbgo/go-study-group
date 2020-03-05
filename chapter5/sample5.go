package chapter5

func SumMulti(x, num int) int {
	var result int
	for i := 0; i < num; i++ {
		result = result + x
	}
	return result
}
