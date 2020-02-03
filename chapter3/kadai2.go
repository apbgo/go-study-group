package chapter3

// 課題2
// Goの言語仕様にはコンストラクタの機能は存在しません。
// しかし慣習的にNew[struct名]の関数を実装してstructの初期化とそのポインタ返却します。
// 以下のstructに対してコンストラクタを実装をしてみましょう。
// コンストラクタ名はNewKadai2, 戻り値はKadai2のポインタにしましょう。
type Kadai2 struct {
	id   int
	name string
}
