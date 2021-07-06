package main

import "github.com/ironprayer/go_crawler_module/filter"

func main() {
	//url := "https://www.pinterest.co.kr/"
	//url := "https://home.kepco.co.kr/kepco/main.do"
	url := []string{"https://naver.com/blog", "https://www.google.co.kr/imgres"}
	//imgURL := "/12.png/test"
	//baseURL := "https://naver.com/"
	//parser.GetContent(url)
	//parser.GetBaseURL(url)
	//parser.GetCleansingURL(baseURL, imgURL)
	filter.RemoveNotValidURLInRobots(url)

}
