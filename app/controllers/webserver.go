package controllers

import (
	"html/template"
	"net/http"
)

// メイン画面の表示
var templates = template.Must(template.ParseFiles("app/views/main.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "main.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/view/", viewHandler)
	return http.ListenAndServe(":8080", nil)
}