package main

import (
	"github.com/ironprayer/go_crawler_module/parser"
)

func main() {
	//url := "https://www.pinterest.co.kr/"
	//url := "https://home.kepco.co.kr/kepco/main.do"
	imgURL := "/12.png/test"
	baseURL := "https://naver.com/"
	//parser.GetContent(url)
	parser.GetCleansingURL(baseURL, imgURL)

}
