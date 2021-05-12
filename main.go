package main

import (
	"github.com/ironprayer/go_crawler_module/parser"
)

func main() {
	url := "https://www.pinterest.co.kr/"
	parser.GetContent(url)
}
