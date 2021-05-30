package parser

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractData struct {
	title       string
	description string
	tags        []string
	url         []string
	//데이터 정의 및 추가해야 함
}

type ImgInfo struct {
	url string
	alt string
}

// error -> 벤인것과 그냥 안되는 페이지 -> 벤 인건 따로 구분해서 시간 조정하는 함수 -> 특정시간이 지나면 벤인건 다시 때린다
// 크롤러라는 게 일정주기로 다시 전체 사이트 접속해봐야겠지. (사이트 꺼졌다던가, 뭐했다던가, 벤 당한건) reuqests 보내는 시간을 조정해줘야됨
// 주기 정하는 것

func GetContent(html string) {
	fmt.Println("Parser Start...")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		log.Fatal(err)
	}

	title := getTitleInDocument(doc)
	description := getDescriptionInDocument(doc)
	imgInfos := getImgDatasInDocument(doc)
	links := getURLsInDocument(doc)

	fmt.Println(title, description, imgInfos, links)

}

func getTitleInDocument(doc *goquery.Document) string {
	docTitle := doc.Find("title")
	return docTitle.Text()
}

func getDescriptionInDocument(doc *goquery.Document) string {
	docMetas := doc.Find("meta")
	description := ""

	docMetas.Each(func(i int, meta *goquery.Selection) {
		if name, _ := meta.Attr("name"); name == "description" {
			description, _ = meta.Attr("content")
		}
	})

	return description
}

func getImgDatasInDocument(doc *goquery.Document) []ImgInfo {
	var imgInfos []ImgInfo
	docImgs := doc.Find("img")
	docPicture := doc.Find("picture")

	// 함수로 만들어서 1개로 줄일 거 생각해야 함.
	docImgs.Each(func(i int, img *goquery.Selection) {
		urlOfImg, _ := img.Attr("src")
		altOfImg, _ := img.Attr("alt")
		imgInfos = append(imgInfos, ImgInfo{urlOfImg, altOfImg})
	})

	docPicture.Each(func(i int, pic *goquery.Selection) {
		urlOfImg, _ := pic.Attr("src")
		altOfImg, _ := pic.Attr("alt")
		imgInfos = append(imgInfos, ImgInfo{urlOfImg, altOfImg})
	})

	return imgInfos
}

// rel = "nofollow" or rel="ugc" 속성 확인 필요
func getURLsInDocument(doc *goquery.Document) []string {
	var URLs []string
	docATags := doc.Find("a")

	docATags.Each(func(i int, tag *goquery.Selection) {
		tempLink, _ := tag.Attr("href")
		link := getCleansingURL(tempLink)

		if link != "" {
			URLs = append(URLs, link)
		}
	})

	return URLs
}

/*
*	작성 필요
*	Case 1 : URL 형식이 아닌 문자열 삭제 (예 : void(0))
*	Case 2 : 기본 URL이 없는 문자열 정제 (예 : /img/example.png) => (base_url)http://demo.samplePage + (extract_url)/img/example.png
*	Case 3 : 정상적인 URL 문자열 (예 : http://demo.samplePage/img/example.png)
 */
func getCleansingURL(url string) string {
	return url
}

/*
func getJsonInDocument(doc *goquery.Documnet) string{
	doc.Find("script").Each(func(i int, script *goquery.Selection) {
		scriptType, _ := script.Attr("type")
		if scriptType == "application/ld+json" {
			fmt.Println(script.Text())
		}
	})
}
*/
