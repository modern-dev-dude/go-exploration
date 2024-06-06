package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func main() {
	bookUrl := "http://localhost:3000/api/book"
	contentType := "application/json"
	body := []byte(`{"bookId": 1}`)

	bookRequest, err := http.NewRequest("POST", bookUrl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	bookRequest.Header.Add("Content-Type", contentType)

	client := &http.Client{}

	res, err := client.Do(bookRequest)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	book := &Book{}
	decoder := json.NewDecoder(res.Body).Decode(book)

	if decoder != nil {
		log.Fatal(err)
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	fmt.Println("Id: ", book.Id)
	fmt.Println("Title: ", book.Title)
	fmt.Println("Author: ", book.Author)
	fmt.Println("Description: ", book.Description)
}
