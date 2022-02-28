package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

// メイン画面の表示
var templates = template.Must(template.ParseFiles("app/views/main.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "main.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// レスポンスで返す用のJSON
type Newline struct {
	// Status int  `json:"status"`
	Text string `json:"text"`
}

// 改行を行う
func newLineHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request arrived")
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

	// 構造体をjson化
	jsondata, err := json.Marshal(newLine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// headerのセット
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.WriteHeader(http.StatusOK)

	// bodyのセット
	w.Write(jsondata)
	// fmt.Fprint(w, string(jsondata))
}

// test
// func testHandler(w http.ResponseWriter, r *http.Request) {
// 	//構造体をjson化
// 	newLine := Newline{"text"}
// 	jsondata, err := json.Marshal(newLine)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// headerのセット
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	// w.WriteHeader(http.StatusOK)

// 	// bodyのセット
// 	w.Write(jsondata)
// 	// fmt.Fprint(w, string(jsondata))
// }

func StartWebServer() error {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/newline/", newLineHandler)

	// cssやjsを読め込めるようにする
	dir, _ := os.Getwd()
	// ディレクトリを指定して公開
	http.Handle("/app/views/", http.StripPrefix("/app/views/", http.FileServer(http.Dir(dir+"/app/views/"))))


	return http.ListenAndServe(":8080", nil)
}