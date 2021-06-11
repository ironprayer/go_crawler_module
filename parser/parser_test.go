package parser_test

import (
	"testing"

	"github.com/ironprayer/go_crawler_module/parser"
	"github.com/stretchr/testify/assert"
)

func TestGetCleansingURL(t *testing.T) {
	actual := parser.GetCleansingURL("http://www.example", "/test.png")
	expected := "http://www.example/test.png"
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다")

}

func TestIsValidURL(t *testing.T) {
	actual := parser.IsValidURL("/test.png")
	expected := true
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다")

	actual = parser.IsValidURL("void(0)")
	expected = false
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다")
}

func TestWriteInFile(t *testing.T) {
	var testList = []string{"File Write Test", "File Write Test 02", "File Write Test 03"}
	parser.WriteLinkInFile(testList)
}

func TestWriteDataInFile(t *testing.T) {
	//오류 확인 필요
	var testList = []parser.ImgInfo{parser.ImgInfo{"File Write Test", "test1"}, parser.ImgInfo{"File Write Test 02", "test2"}, parser.ImgInfo{"File Write Test 03", "test3"}}
	parser.WriteImgInFile(testList)
}
