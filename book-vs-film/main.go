package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io"
	"encoding/json"
)

type VolumeInfo struct {
	Title string `json:"title"`
	Authors []string `json: "authors"`
	Language string `json:"language"`
	PublishedDate string `json:"publishedDate"`
	Rating float32 `json:"averageRating"`
	RatingsCount int `json:"ratingsCount"`
	InfoLink string `json:"infoLink"`
}

type VolumeItem struct {
	Kind string `json:"kind"`
	Id string `json:"id"`
	SelfLink string `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type BooksAPIResponse struct {
	Kind string `json:"kind"`
	TotalItems int `json:"totalItems"`
	Items []VolumeItem `json:"items"`
}

type BookInfo struct {
	Kind string `json:"kind"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

func filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, element := range slice {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}


var books_api_key = "AIzaSyB9NE_zsBc-q99c6omt6aMuWArx-RHcIzc"

func PrettyPrint(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b))
}

func GetBookInfo(title string, author string) (BookInfo, error) {
	base_url := "https://www.googleapis.com/books/v1/volumes"
	parsed_url, err := url.Parse(base_url)
	params := url.Values{}
	params.Add("q", title)
	params.Add("key", books_api_key)
	params.Add("langRestrict", "en")       // Only english
	params.Add("printType", "books")       // Avoids magazines
	params.Add("projection", "full")       // Requests full metadata

	parsed_url.RawQuery = params.Encode()
	response, err := http.Get(parsed_url.String())
	if err != nil {
		return BookInfo{}, err
	}
	body, err := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return BookInfo{}, fmt.Errorf("Response status code is %d when searching for books.\n%s\n", response.StatusCode, string(body))
	}
	if err != nil {
		return BookInfo{}, err
	}
	var found_books BooksAPIResponse 
	err = json.Unmarshal(body, &found_books)
	fmt.Printf("Searching for author: %s\nAuthors:\n", author)
	books := filter(found_books.Items, func (book VolumeItem) bool {
		return true
	})
	for _, book := range books {
		PrettyPrint(book)
	}
	return BookInfo{}, nil
	/*
	if len(books) == 0 {
		return BookInfo{}, fmt.Errorf("No books found by title %s\n", title)
	}
	book := books[0]
	base_url = "https://www.googleapis.com/books/v1/volumes/" + book.Id
	parsed_url, err = url.Parse(base_url)
	params = url.Values{}
	params.Add("key", books_api_key)
	parsed_url.RawQuery = params.Encode()
	response, err = http.Get(parsed_url.String())
	if err != nil {
		return BookInfo{}, err
	}
	body, err = io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return BookInfo{}, fmt.Errorf("Response status code is %d when getting book info.\n%s\n", response.StatusCode, string(body))
	}
	if err != nil {
		return BookInfo{}, err
	}
	var book_info BookInfo
	err = json.Unmarshal(body, &book_info)
	return book_info, err
	*/
}

func main() {
	book, err := GetBookInfo("Хоббит", "J.R.R. Tolkien")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Found Book:", book)
}


