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