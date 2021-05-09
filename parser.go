package parser

import (
	"fmt"
	"net/http"
)

type extractData struct {
	title       string
	description string
	//데이터 정의 및 추가해야 함
}

func getContent(pageURL string) {
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	//Error Check 필요

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	//Error Check 필요

	//extractData 정보 추출
	docTitle := doc.Find("title")
	docDescription := doc.Find("meta")

	extractData := extractData{title: docTitle, description: docDescription}

	fmt.Println(extractData)
}
