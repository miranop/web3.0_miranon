package main

import (
	"fmt"
	"net/http"
)

func main() {
	//HTMLファイルを提供するルートディレクトリを指定する
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("ローカルホスト8080を開けました")
	http.ListenAndServe(":8080", nil)
}
