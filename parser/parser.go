package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type extractData struct {
	title       string
	description string
	//데이터 정의 및 추가해야 함
}

func GetContent(pageURL string) {
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	//Error Check 필요
	fmt.Println(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	//Error Check 필요

	/*extractData 정보 추출
	google SCO 참고

	구조화된 데이터
	tags : <script type="application/ld+json"></script>

	이미지
	tags : <img> or <picture>
	attr : Alt

	설명
	tags : meta
	attr : name : description

	*/
	docTitle := doc.Find("title")
	docMetas := doc.Find("meta")
	docImgs := doc.Find("img")
	docPictures := doc.Find("piture")
	docATags := doc.Find("a")

	fmt.Println(docTitle)
	fmt.Println((*docTitle).Text())

	fmt.Println(docMetas.Text())
	fmt.Println("Atags", docATags.Length())

	/* 빈 공백만 출력됨
	docMetas.Each(func(i int, meta *goquery.Selection) {
		fmt.Println(meta.Text())
	})
	*/
	fmt.Println(docMetas)

	fmt.Println(docTitle.Length(), docMetas.Length(), docImgs.Length(), docPictures.Length())

	//extractData := extractData{title: docTitle, docMetas, docImgs, docPictures}

	//fmt.Println(extractData)
}

// 페이지에 존재하는 URL 링크 가져오기
// rel = "nofollow" or rel="ugc" 속성 확인 필요
/*
func getUrls(doc string) string {
	aTags := doc.Find("a")

	return aTags
}
*/
