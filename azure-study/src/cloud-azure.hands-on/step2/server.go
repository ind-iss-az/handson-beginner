package main

import (
	"log"
	"net/http"
	"html/template"
)

func HelloHandler (writer http.ResponseWriter, request *http.Request) {
	// テンプレートのパース処理を行う
	tmpl, _ := template.ParseFiles("tmpl-step2.html")

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

// Add Step2 HealthCheck
func HealthCheackHandler (writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("Status 200 OK!"))
}

func main() {
	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/hello", HelloHandler)
	// Add Step2
	http.HandleFunc("/healthcheck", HealthCheackHandler)

	server.ListenAndServe()
}