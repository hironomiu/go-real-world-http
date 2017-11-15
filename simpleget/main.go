package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
	log.Println("Content-Length:", resp.Header["Content-Length"])
	log.Println("Content-Type:", resp.Header.Get("Content-Type"))
	log.Println("Content-Type:", resp.Header["Content-Type"])
	log.Println(string(body))
}
