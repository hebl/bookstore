package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hebl/bookstore"
)

func main() {

	var isbn = flag.String("i", "9787115385376", "BOOK ISBN number")
	flag.Parse()
	url := fmt.Sprintf("https://api.douban.com/v2/book/isbn/%s", *isbn)

	req, err := http.NewRequest("GET", url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var book *bookstore.BookInfo
	err = json.Unmarshal(body, &book)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ISBN    :\t%s\n", book.ISBN13)
	fmt.Printf("名    称:\t%s\n", book.Title)
	fmt.Printf("名    称:\t%s\n", book.SubTitle)
	fmt.Printf("原    名:\t%s\n", book.OriginTitle)
	fmt.Printf("书    系:\t%s\n", book.Series.Title)
	fmt.Printf("作    者:\t%v\n", strings.Join(book.Author, ", "))
	fmt.Printf("译    者:\t%v\n", strings.Join(book.Translator, ", "))
	fmt.Printf("出 版 社:\t%v\n", book.Publisher)
	fmt.Printf("出版时间:\t%v\n", book.Pubdate)
	fmt.Printf("价    格:\t%v\n", book.Price)
	fmt.Printf("页    数:\t%v\n", book.Pages)
	fmt.Printf("图    片:\t%v\n", book.Image)
}
