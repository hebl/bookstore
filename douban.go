package bookstore

import "encoding/json"

/**
参考 http://developers.douban.com/wiki/?title=book_v2
*/

// BookInfo 图书的基本信息
type BookInfo struct {
	ID          string `json:"id"`
	ISBN10      string `json:"isbn10"`
	ISBN13      string `json:"isbn13"`
	Title       string `json:"title"`
	OriginTitle string `json:"origin_title"`
	AltTitle    string `json:"alt_title"`
	SubTitle    string `json:"subtitle"`
	URL         string `json:"url"`
	Alt         string `json:"alt"`
	Image       string `json:"image"`
	Images      struct {
		Small  string `json:"small"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
	} `json:"images"`
	Author     []string `json:"author"`
	Translator []string `json:"translator"`
	Publisher  string   `json:"publisher"`
	Pubdate    string   `json:"pubdate"`
	Rating     struct {
		Max       int    `json:"max"`
		NumRaters int    `json:"numRaters"`
		Average   string `json:"average"`
		Min       int    `json:"min"`
	} `json:"rating"`
	Tags []struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
	} `json:"tags"`
	Binding string `json:"binding"`
	Price   string `json:"price"`
	Series  struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"series"`
	Pages       string `json:"pages"`
	AuthorIntro string `json:"author_intro"`
	Summary     string `json:"summary"`
	Catalog     string `json:"catalog"`
	EBookURL    string `json:"ebook_url"`
	EBookPrice  string `json:"ebook_price"`
}

// Unmarshal 解析
func Unmarshal(body []byte) (*BookInfo, error) {
	r := &BookInfo{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil
}
