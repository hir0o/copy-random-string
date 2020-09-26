package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"

	"github.com/atotto/clipboard"
)

func main() {
	// オプションを取得
	length := flag.Int("l", 6, "文字の長さ")
	flag.Parse()
	// ランダム文字列を作成
	random, _ := MakeRandomStr(uint32(*length))
	// クリップボードにコピー
	clipboard.WriteAll(random)
	// 結果を出力
	fmt.Println(random + "をクリップボードにコピーしました。")
}

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
