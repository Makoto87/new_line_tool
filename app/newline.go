package app

import (
	"regexp"
)

// リクエスト・レスポンス用の構造体
type Newline struct {
	Text string `json:"text"`
}

// テキストを改行させる
func (n *Newline) MakeNewline() {
	// 正規表現でスペースを改行に変える
	// 正規表現以外のやり方はないのか？
	var findSpace = regexp.MustCompile("( |　)+")
	replacedSentence := findSpace.ReplaceAllString(n.Text, "\n")

	// replaceした文字列をNewlineにセット
	n.Text = replacedSentence
}
