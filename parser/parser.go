package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type extractData struct {
	title       string
	description string
	tags        []string
	url         []string
	//데이터 정의 및 추가해야 함
}

// error -> 벤인것과 그냥 안되는 페이지 -> 벤 인건 따로 구분해서 시간 조정하는 함수 -> 특정시간이 지나면 벤인건 다시 때린다
// 크롤러라는 게 일정주기로 다시 전체 사이트 접속해봐야겠지. (사이트 꺼졌다던가, 뭐했다던가, 벤 당한건) reuqests 보내는 시간을 조정해줘야됨
// 주기 정하는 것

/*
func GetDocument(pageURL string) (*goquery.Document, error) {
	res, err := http.Get(pageURL)
	if err != nil {
		return nil, nil
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	return doc, err
}
*/

func GetContent(pageURL string) {
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	//Error Check 필요
	fmt.Println(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	fmt.Println("error")
	fmt.Println(err)
	if err == nil {
		fmt.Println("Not Err")
	}
	//Error Check 필요

	/*extractData 정보 추출
	google SEO 참고 & Pinterest = Img

	제목
	tags : title

	구조화된 데이터
	tags : <script type="application/ld+json"></script>

	이미지
	tags : <img> or <picture>
	attr : alt

	설명
	tags : meta
	attr : name : description

	*/
	docTitle := doc.Find("title")
	docMetas := doc.Find("meta")
	//개수가 많다. Url : img, pic, a    1 : n => 어떤식으로 데이터를 떨궈줘야 할지 /
	//img1 : 원천 데이터
	docImgs := doc.Find("img")
	docPictures := doc.Find("picture")
	//docATags := doc.Find("a")

	fmt.Println(docTitle.Text())
	fmt.Println(docMetas.Length())
	//fmt.Println(docMetas.Attr("name"))
	//fmt.Println((*docMetas).Attr("name"))

	/*
		docMetas.Each(func(i int, meta *goquery.Selection) {
			fmt.Println(meta.Attr("name")) //meta 전체 이름 출력
			if name, _ := meta.Attr("name"); name == "description" {
				description, _ := meta.Attr("content")
				fmt.Printf("Doc Description: %s\n", description)
			}
		})
	*/
	/*
		docImgs.Each(func(i int, img *goquery.Selection) {
			fmt.Println(img.Attr("alt"))
		})
	*/

	/*
		fmt.Printf("A Link Count : %d\n", docATags.Length())
		docATags.Each(func(i int, tag *goquery.Selection) {
			link, _ := tag.Attr("href")
			fmt.Printf("No%d : %s\n", i, link)
		})
	*/

	doc.Find("script").Each(func(i int, script *goquery.Selection) {
		scriptType, _ := script.Attr("type")
		if scriptType == "application/ld+json" {
			fmt.Println(script.Text())
		}
	})

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
