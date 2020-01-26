package main

import (
	"fmt"
	"github.com/go-numb/go-trade-utility/plot"
	"math/rand"
	"time"
)

func main() {
	// 時間によって変化する整数を設定する
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// 本当にランダムに整数が生成されているのか確認してみる
	// 数が少なくてわからなかったので多くしてみる
	count := 10000
	for i := 0; i < count; i++ {
		fmt.Printf("%+v\n", r.Int63())
	}
	// 多分乱数が生成された
	// が、もっと確認してみる
	// 出現数をメモリに保存する
	fx := make([]float64, count)
	// fxの参照数も作っとく（グラフのため、メモリみたいなもの
	index := make([]float64, count)
	for i := 0; i < count; i++ {
		// appendは配列長がわからないまま、追加場所を探しながら追加するので、極力避けたい
		// 長さを教えて、追加する場所を教えてあげる

		// count数乱数を作ってるんだから、count長になる
		fx[i] = float64(r.Int63())
		index[i] = float64(i)
	}

	// 図で確認してみる
	scatter := plot.NewScatter("random", "n", "random_n", index, fx)
	scatter.Title = "random numbers"

	scatter.Save("./plot.png")

	// ランダムだよね
	// ランダムっていうのは、常々まばらなことであって、等間隔のことではないことを理解しておきたい

	// 乱数を何に使うのか？
	// - 乱数の余りで挙動を分けたい
	// - 7で割れば、割れる場合1・割れない場合6で７パターン出現するはず
	for i := range fx {
		n := int(fx[i]) % 7
		switch n {
		case 0:
			fmt.Printf("%+v - 割れたでしょ\n", n)

		default: // 割れていない場合の処理
			fmt.Printf("余り = %d\n", n)
			switch n {
			case 1:
			case 2: // ... などすれば、７パターンの状態分けができる。状態一つの確率は0 < n <= ∞ならば、1/7である
			}
		}
	}
}
