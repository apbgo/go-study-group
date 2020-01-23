package chapter1

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/apbgo/go-study-group/chapter1/lib"
)

// Calc opには+,-,×,÷の4つが渡ってくることを想定してxとyについて計算して返却(正常時はerrorはnilでよい)
// 想定していないopが渡って来た時には0とerrorを返却
func CalcAns(op string, x, y int) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "×":
		return x * y, nil
	case "÷":
		return x / y, nil
	default:
		return 0, errors.New("opが不正です")
	}
}

// StringEncode 引数strの長さが5以下の時キャメルケースにして返却、それ以外であればスネークケースにして返却
func StringEncodeAns(str string) string {
	if len(str) <= 5 {
		return lib.ToCamel(str)
	}
	return lib.ToSnake(str)
}

// Sqrt 数値xが与えられたときにz²が最もxに近い数値zを返却
func SqrtAns1(x float64) float64 {
	z := 1.0
	// 10回繰り返す例
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SqrtAns2(x float64) float64 {
	z := 1.0
	// 差分が小さくなるまで繰り返す例
	for d := 1.0; d*d > 1e-10; z -= d {
		d = (z*z - x) / (2 * z)
	}
	return z
}

// Pyramid x段のピラミッドを文字列にして返却
// 期待する戻り値の例：x=5のとき "1\n12\n123\n1234\n12345"
// （x<=0の時は"error"を返却）
func PyramidAns(x int) string {
	// 素直にStringを連結してもOK（速度は遅い）
	// var str string
	var sb strings.Builder

	for i := 0; i < x; i++ {
		if i != 0 {
			sb.WriteString("\n")
			// str += "\n"
		}
		for j := 0; j < i+1; j++ {
			sb.WriteString(strconv.Itoa(j + 1))
			// str += strconv.Itoa(j + 1)
		}
	}
	return sb.String()
	//return str
}

// StringSum x,yをintにキャストし合計値を返却 (正常終了時、errorはnilでよい)
// キャスト時にエラーがあれば0とエラーを返却
func StringSumAns(x, y string) (int, error) {
	sx, err := strconv.Atoi(x)
	if err != nil {
		return 0, err
	}
	sy, err := strconv.Atoi(y)
	if err != nil {
		return 0, err
	}
	return sx + sy, nil
}

// SumFromFileNumber ファイルを開いてそこに記載のある数字の和を返却
func SumFromFileNumberAns(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		sum += num
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}
