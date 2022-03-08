package app

import (
	"embed"
	"encoding/json"
	"io"

	"html/template"
	"net/http"
)

// html, css. jsを埋め込みする
//go:embed views/*
var views embed.FS

// メイン画面の表示
var templates = template.Must(template.ParseFS(views, "views/main.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "main.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// リクエストをもらい、レスポンスを返す
func newLineHandler(w http.ResponseWriter, r *http.Request) {
	// POSTじゃなかったらエラー
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// josnじゃなかったらエラー
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Bodyを読み込み
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// jsonを読み込み
	// decoderとどっちを使えばいい？
	var newLine Newline
	if err = json.Unmarshal(body, &newLine); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// decoerの場合
	// json.NewDecoder(r.Body).Decode(&newLine)

	// 文字列内のスペースを改行に変換
	newLine.MakeNewline()

	// 構造体をjson化
	jsondata, err := json.Marshal(newLine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// headerのセット
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// bodyのセット
	w.Write(jsondata)
}

// httpサーバーを立ち上げ
func StartWebServer() error {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/newline/", newLineHandler)

	// go:embedを利用し、viewsディレクトリ以下にあるファイルを全て公開
	http.Handle("/views/", http.FileServer(http.FS(views)))

	return http.ListenAndServe(":8080", nil)
}
