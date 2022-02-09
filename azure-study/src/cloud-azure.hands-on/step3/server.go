package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

var flg_init = false

func HelloHandler (writer http.ResponseWriter, request *http.Request) {
	// テンプレートのパース処理を行う
	tmpl, _ := template.ParseFiles("tmpl-step3.html")

	// テンプレートに埋め込むデータの作成を行う
	date := GetDatetime()
	hostname := GetHostname()

	var db_data string = ""
	if flg_init != false {
		db_data = fmt.Sprintf("%+v",PutMessages(date, hostname))
	}


    tmplData := map[string]string{
        "Date": date,
		"Hostname": hostname,
		"DBdata": db_data,
	}
	
	// テンプレートを描画処理を行う
    if err := tmpl.Execute(writer, tmplData); err != nil {
        log.Fatal(err)
	}
}

func HealthCheackHandler (writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("Status 200 OK!"))
}

// Add Step3
func init(){
	// Parameter Store -> Load
	paramstore, err := newParameterStore()
	if err != nil {
		log.Println("errorparam")
		//log.Fatal(err)
		return
	}

	// sql.Connect
	dbsetting := newDbSetting(&paramstore.param)
	Db, err = sql.Open("sqlserver", dbsetting.dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	flg_init = true
}

func main() {
	server := http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/healthcheck", HealthCheackHandler)

	server.ListenAndServe()
}