package chapter2

import "fmt"

// 引数のスライスsliceの要素数が
// 0の場合、0とエラー
// 2以下の場合、要素を掛け算
// 3以上の場合、要素を足し算
// を返却。正常終了時、errorはnilでよい
func CalcAns(slice []int) (int, error) {
	// TODO Q1
	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf
	switch len(slice) {
	case 0:
		return 0, fmt.Errorf("slice length is zero")
	case 1:
		return slice[0], nil
	case 2:
		return slice[0] * slice[1], nil
	default:
		var ret int
		for _, v := range slice {
			ret += v
		}
		return ret, nil
	}
}

// 構造体Numberを3つの要素数から成るスライスにして返却
// 3つの要素の中身は[{1} {2} {3}]とし、append関数を使用すること
func NumbersAns() []Number {
	// TODO Q2
	result := make([]Number, 0, 3)

	for i := 1; i <= 3; i++ {
		result = append(result, Number{index: i})
	}

	return result
}

// 引数mをforで回し、「値」部分だけの和を返却
// キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
// キー「yon」に関しては完全一致すること
func CalcMapAns(m map[string]int) int {
	// TODO Q3
	var result int

	_, ok := m["yon"]
	if ok {
		delete(m, "yon")
	}

	for _, v := range m {
		result += v
	}

	return result
}

// 与えられたスライスのModel全てのValueに5を足す破壊的な関数を作成
func AddAns(models []Model) {
	// TODO  Q4
	for i := range models {
		models[i].Value += 5
	}
}

// 引数のスライスには重複な値が格納されているのでユニークな値のスライスに加工して返却
// 順序はスライスに格納されている順番のまま返却すること
// ex) 引数:[]slice{21,21,4,5} 戻り値:[]int{21,4,5}
func UniqueAns(slice []int) []int {
	// TODO Q5

	result := make([]int, 0)
	uniqueMap := make(map[int]struct{})
	for _, v := range slice {
		_, ok := uniqueMap[v]
		if ok {
			// uniqueMapに含まれている
			continue
		}
		result = append(result, v)
		uniqueMap[v] = struct{}{}
	}
	return result
}

// 連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
func FibonacciAns() func() int {
	// TODO Q6 オプション

	a := 0
	b := 1

	return func() int {
		a, b = b, a+b
		return b - a
	}
}
