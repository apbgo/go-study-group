package chapter2

// 引数のスライスsliceの要素数が
// 0の場合、0とエラー
// 2以下の場合、要素を掛け算
// 3以上の場合、要素を足し算
// を返却。正常終了時、errorはnilでよい
func Calc(slice []int) (int, error) {
	// TODO Q1
	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	return 0, nil
}

type Number struct {
	index int
}

// 構造体Numberを3つの要素数なら成るスライスにして返却
// 条件1: 3つの要素の中身は[{1} {2} {3}]とすること
// 条件2: append関数を使用すること
func Numbers() []Number {
	// TODO Q2

	return nil
}

// 引数mをforで回し、「値」部分だけの和を返却
// キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
func CalcMap(m map[string]int) int {
	// TODO Q3

	return 0
}

// 連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
func Fibonacci() func() int {
	// TODO Q4

	return nil
}
