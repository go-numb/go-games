package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"os"
	"time"
)

var programs = []string{
	"異常生物見聞録",
	"アイカツオンパレード",
	"アイドルマスターシンデレラガールズ",
	"アサシンブライド",
	"アズールレーン",
	"鬼滅の刃",
	"兄に付ける薬はない！",
	"あひるの空",
	"アフリカのサラリーマン",
	"雨色ココア",
	"洗い屋さん",
	"荒ぶる季節の乙女どもよ。",
	"異世界かるてっと",
	"上野は不器用",
	"炎炎ノ消防隊",
	"俺を好きなのはお前だけかよ",
	"カケグルイ",
	"からかい上手の高木さん",
	"歌舞伎町シャーロック",
	"ケムリクサ",
	"賢者の孫",
	"五等分の花嫁",
	"PSYCHO-PASS3",
	"食戟のソーマ",
	"女子高生の無駄つかい",
	"世話やき狐の仙狐さん",
	"川柳少女",
	"盾の勇者の成り上がり",
	"ちはやふる3",
	"ダンベル何キロ持てる？",
	"ダンジョンに出会いを求めるのは間違っているだろうか",
	"どろろ",
	"Dr.Stone",
	"七つの大罪",
	"ハイスコアガール",
	"僕のヒーローアカデミア",
	"ぼくたちは勉強ができない",
	"ポプテピピック",
}

type Client struct {
	db *sql.DB
}

func New() *Client {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return nil
	}

	_, err = tx.Query("select * from programs;")
	if err != nil {
		fmt.Println("選択するdatabase tableがないので作る")
		db.Exec("create table programs (id primary key, value text);")
		state, _ := tx.Prepare("insert into programs (id, value) values (?, ?);")
		for i := range programs {
			state.Exec(i, programs[i])
		}
		_, err = tx.Query("select * from programs;")
		if err != nil {
			fmt.Println("問題用database tableをつくったが、それでもエラー", err.Error())
			return nil
		}
	}
	tx.Commit()

	// var count int
	// for rows.Next() {
	// 	count++
	// }

	// if count == 0 { // countが0ならば、初期問題を保存する
	// 	tx.Exec("create table programs (id primary key, value text);")
	// 	state, _ := tx.Prepare("insert into programs values (?, ?);")
	// 	for i := range programs {
	// 		state.Exec(i, programs[i])
	// 	}
	// }

	return &Client{
		db: db,
	}
}

func (p *Client) Close() error {
	if err := p.db.Close(); err != nil {
		return err
	}

	return nil
}

func main() {
	// 終了・途中終了を検出
	done := make(chan os.Signal, 1)

	// 時間を背景に、整数の乱数を作る
	s := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s)

	// データベースクライアントをつくる
	client := New()
	defer client.Close()

	// 10問出題中正解率を求める
	var (
		count   = 10
		correct = 0
		retry   = 0
	)
	for i := 0; i < count; i++ {
	RELOAD:
		var result int
		n := random.Intn(1000)
		client.db.QueryRow(fmt.Sprintf("select exists(select * from programs where id == %d);", n)).Scan(&result)
		if result != 1 {
			retry++
			goto RELOAD
		}
		rows, err := client.db.Query(fmt.Sprintf("select value from programs where id == %d;", n))
		if err != nil {
			retry++
			goto RELOAD
		}

		var q string
		for rows.Next() {
			rows.Scan(&q)
			if q != "" {
				fmt.Printf("出題:%d/%d\n%s\n", i+1, count, q)
				break
			}
			retry++
			goto RELOAD
		}

		var input string
		fmt.Scan(&input)

		if input != q { // 出題文と入力文を照らし合わせる
			fmt.Println("✗: 不正解！！")
			continue
		}

		// 正解数を数える
		fmt.Println("○: 正解！！")
		correct++
	}

	fmt.Printf("%d問全問終了！\n正解率: %.1f％\n\n", count, float64(correct)/float64(count)*100)
	fmt.Printf("開発メモ: retry = %d\n", retry)
	time.Sleep(3 * time.Second)

	<-done
}
