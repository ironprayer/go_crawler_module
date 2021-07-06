package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRobotsURL(t *testing.T) {
	actual := getRobotsURL("https://naver.com/blog")
	expected := "https://naver.com/robots.txt"
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다.")

	actual = getRobotsURL("https://www.google.co.kr/search")
	expected = "https://www.google.co.kr/robot.txt"
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다.")
}

func TestGetTargetURL(t *testing.T) {
	actual := getTargetURL("https://www.google.co.kr/search")
	expected := "/search"
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다.")
}

//google group false 나와야 하는데 true가 나옴 확인 필요
func TestIsAllowURL(t *testing.T) {
	url := "https://www.google.co.kr/search/about"
	agent := "*"
	robotsURL := getRobotsURL(url)
	robotsRule := getRobotsRules(robotsURL)
	targetURL := getTargetURL(url)
	isAllowURL := isAllowURLWithRobots(robotsRule, agent, targetURL)
	expected := true

	assert.Equal(t, expected, isAllowURL, "기대값과 결과값이 다릅니다.")
}
