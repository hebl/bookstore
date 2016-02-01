package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hebl/bookstore"
)

func main() {

	var isbn = flag.String("i", "9787115385376", "BOOK ISBN number")
	var list = flag.Bool("l", false, "列表显示")
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
	if *list {
		fmt.Printf("ISBN    :\t%s\n", book.ISBN13)
		fmt.Printf("名    称:\t%s\n", book.Title)
		fmt.Printf("作    者:\t%v\n", book.Author)
		fmt.Printf("出 版 社:\t%v\n", book.Publisher)
		fmt.Printf("出版时间:\t%v\n", book.Pubdate)
		fmt.Printf("价    格:\t%v\n", book.Price)
	}

}
