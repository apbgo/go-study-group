package chapter5

import (
	"testing"
)

func BenchmarkSum(b *testing.B) {
	var result int
	b.ResetTimer() // ここから計測

	for i := 0; i < b.N; i++ {
		result++
	}
}

func BenchmarkSumMulti(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SumMulti(1, 100)
	}
}

/// メモリアロケーションについてのベンチマーク -----------------------

const maxNum = 100

func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base := []string{}
		for j := 0; j < maxNum; j++ {
			// 都度append
			base = append(base, "hoge")
		}
	}
}

func BenchmarkAppend_AllocateOnce(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		//最初に長さを決める
		base := make([]string, 0, maxNum)
		for j := 0; j < maxNum; j++ {
			base = append(base, "hoge")
		}
	}
}

/// 並列処理についてのベンチマーク -----------------------

func BenchmarkSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMulti(1, 10)
	}
}

func BenchmarkParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SumMulti(1, 10)
		}
	})
}
