package filter

import (
	"log"
	"net/http"
	"strings"

	"github.com/temoto/robotstxt"
)

func removeDuplicateURL(urls []string) {
	// 중복 제거 확인 해야할듯

}
func removeVisitedURL(urls []string) {
	// 파일 또는 DB에서 가져와야할 것 같은데
}
func RemoveNotValidURLInRobots(urls []string) []string {

	agent := "*"
	validURLs := make([]string, 0)

	for _, url := range urls {
		robotsURL := getRobotsURL(url)
		targetURL := getTargetURL(url)
		robotsRule := getRobotsRules(robotsURL)
		isAllowURL := isAllowURLWithRobots(robotsRule, agent, targetURL)

		if isAllowURL == true {
			validURLs = append(validURLs, url)
		}
		//fmt.Println(url, robotsURL, robotsRule, isAllowURL, agent)
	}
	return validURLs
}

func getRobotsURL(url string) string {
	urlUnits := strings.Split(url, "/")
	baseURLUnits := urlUnits[:3]
	baseURL := strings.Join(baseURLUnits, "/")
	robotsFileName := "robots.txt"
	robotsURL := baseURL + "/" + robotsFileName

	return robotsURL
}

func getTargetURL(url string) string {
	urlUnits := strings.Split(url, "/")
	targetURLUnits := urlUnits[3:]
	targetURL := "/" + strings.Join(targetURLUnits, "/")

	return targetURL
}

func getRobotsRules(robotsURL string) *robotstxt.RobotsData {
	res, err := http.Get(robotsURL)
	errCheck(err)
	defer res.Body.Close()

	robots, err := robotstxt.FromResponse(res)
	errCheck(err)

	return robots
}

func isAllowURLWithRobots(robots *robotstxt.RobotsData, agent string, targetURL string) bool {
	allow := robots.TestAgent(targetURL, "*")
	return allow
}

func errCheck(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
