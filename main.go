package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Category ...
type Category struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Color string `json:"color"`
}

// Article ...
type Article struct {
	Category        `json:"category"`
	Likes           int    `json:"likes"`
	Responses       []int  `json:"responses"`
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	CreatedDateTime string `json:"createdDateTime"`
	Slug            string `json:"slug"`
}

type logWriter struct {
	Status   bool `json:"status"`
	Articles []Article
}

func main() {
	resp, err := http.Get("https://dponomar.dev/mock/articles.json")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(&lw, resp.Body)
	fmt.Println("Amount of articles at dponomar.dev is:", len(lw.Articles))
}

func (lw *logWriter) Write(bs []byte) (int, error) {
	json.Unmarshal(bs, &lw)
	return len(bs), nil
}
