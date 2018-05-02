package util

import (
	"math/rand"
	"time"
)

// GetRand ... 与えられた引数を最大値とする乱数を返す
func GetRand(d int) int {
	// デフォルトの擬似乱数生成機のシードを作成
	// Seed内に固定の値を設定してしまうと固定値を毎回返すので、時間等毎回変化する値を与えるとGood!
	rand.Seed(time.Now().UnixNano())
	return int(rand.Float64() * float64(d))
}
