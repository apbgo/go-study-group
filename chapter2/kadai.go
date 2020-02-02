package chapter2

import (
	"fmt"
)

// 引数のスライスsliceの要素数が
// 0の場合、0とエラー
// 2以下の場合、要素を掛け算
// 3以上の場合、要素を足し算
// を返却。正常終了時、errorはnilでよい
func Calc(slice []int) (int, error) {
	l := len(slice)
	switch l {
	case 0 :
		return 0, fmt.Errorf("sliceのサイズが0です")
	case 1 :
		return slice[0], nil
	case 2 :
		return slice[0] * slice[1], nil
	default:
		var sum int
		for i := range slice {
			sum += slice[i]
		}
		return sum, nil
	}
}

type Number struct {
	index int
}

// 構造体Numberを3つの要素数から成るスライスにして返却
// 3つの要素の中身は[{1} {2} {3}]とし、append関数を使用すること
func Numbers() []Number {
	nums := make([]Number, 0, 3)
	nums = append(nums, Number{index:1})
	nums = append(nums, Number{index:2})
	nums = append(nums, Number{index:3})
	return nums
}

// 引数mをforで回し、「値」部分だけの和を返却
// キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
func CalcMap(m map[string]int) int {
	var sum int
	for key, val := range m {
		if key == "yon" {
			continue
		}
		sum += val
	}
	return sum
}

type Model struct {
	Value int
}

// 与えられたスライスのModel全てのValueに5を足す破壊的な関数を作成
func Add(models []Model) {
	for i := range models {
		models[i].Value += 5
	}
}

// 引数のスライスには重複な値が格納されているのでユニークな値のスライスに加工して返却
// 順序はスライスに格納されている順番のまま返却すること
// ex) 引数:[]slice{21,21,4,5} 戻り値:[]int{21,4,5}
func Unique(slice []int) []int {
	uniqueMap := make(map[int]struct{})
	result := make([]int, 0, len(slice))

	for _, val := range slice {
		if _, ok := uniqueMap[val]; ok {
			continue
		}
		uniqueMap[val] = struct{}{}
		result = append(result, val)
	}
	return result
}

// 連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
func Fibonacci() func() int {
	var (
		ex1 = 0
		ex2 = 1
	)
	return func() int {
		ex1, ex2 = ex1+ex2, ex1
		return ex2
	}
}
