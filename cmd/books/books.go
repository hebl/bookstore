package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hebl/bookstore"
)

/**
 * 根据输入一个含有书ISBN号的文件，输出书的详细信息
 */

func pull(isbn string) (*bookstore.BookInfo, error) {
	url := fmt.Sprintf("https://api.douban.com/v2/book/isbn/%s", isbn)

	req, err := http.NewRequest("GET", url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error in request %s: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error in read body: %v\n%s", err, string(body))
		return nil, err
	}

	var book *bookstore.BookInfo
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Fatalf("Error in Unmarshal: %v\n%s", err, string(body))
		return nil, err
	}

	return book, nil
}

// ISBN10 ISBN13 Title SubTitle Author Translator Publisher Pubdate Price Series Pages
func prettyBook(b *bookstore.BookInfo) {
	fmt.Printf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n", b.ISBN10, b.ISBN13, b.Title, b.SubTitle,
		b.OriginTitle, sa(b.Author), sa(b.Translator), b.Publisher, b.Pubdate, y(b.Price),
		b.Series.Title, b.Pages, b.Image)
}

func sa(a []string) string {
	return strings.Join(a, ", ")
}

func y(yuan string) string {

	return strings.TrimSuffix(yuan, "元")
}

func main() {
	var filename = flag.String("i", "book.txt", "BOOK ISBN list file")
	flag.Parse()

	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("Error read file %s: %s", *filename, err)
		return
	}
	list := strings.Split(string(data), "\n")

	for _, isbn := range list {
		if len(isbn) > 0 {
			book, err := pull(isbn)
			if err != nil {
				log.Fatalf("ERROR %s %v", isbn, err)
			}
			prettyBook(book)
		}

	}
}
