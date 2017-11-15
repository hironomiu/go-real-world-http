package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	//fileWriter, err := writer.CreateFormFile("thumbail", "photo.jpg")
	//if err != nil {
	//	panic(err)
	//}

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data;name="thumbnail";filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("/Users/miurahironori/go/src/real-world-http/multipart/photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
