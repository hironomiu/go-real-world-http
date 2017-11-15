package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("/Users/miurahironori/go/src/real-world-http/post/main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)

	reader := strings.NewReader("テキスト")
	hoge, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status", hoge.Status)
}
