package main

import (
	"os"
	"log"
	"time"
	"net/http"
	"html/template"
)

func GetDatetime() string {
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	const layout = "2006-01-02 15:04:05"
	return nowJST.Format(layout)
}

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func HelloHandler (writer http.ResponseWriter, request *http.Request) {
	// テンプレートのパース処理を行う
	tmpl, _ := template.ParseFiles("tmpl-step1.html")

	// テンプレートに埋め込むデータの作成を行う
    tmplData := map[string]string{
        "Date": GetDatetime(),
		"Hostname": GetHostname(),
		"Message": "Hello World!",
    }

	// テンプレートを描画処理を行う
    if err := tmpl.Execute(writer, tmplData); err != nil {
        log.Fatal(err)
    }
}

func main() {
	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/hello", HelloHandler)
	
	server.ListenAndServe()
}