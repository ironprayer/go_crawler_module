package filter

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/temoto/robotstxt"
)

func removeDuplicateURL(urls []string) {}
func removeVisitedURL(urls []string)   {}
func RemoveNotValidURLInRobots(urls []string) {

	agent := "*"

	for _, url := range urls {
		robotsURL := getRobotsURL(url)
		robotsRule := getRobotsRules(robotsURL)
		isAllowURL := isAllowURLWithRobots(robotsRule, agent, url)
		fmt.Println(url, robotsURL, robotsRule, isAllowURL, agent)
	}
}

func getRobotsURL(url string) string {
	urlUnits := strings.Split(url, "/")
	baseURLUnits := urlUnits[:3]
	baseURL := strings.Join(baseURLUnits, "/")
	robotsFileName := "robots.txt"
	robotsURL := baseURL + "/" + robotsFileName

	return robotsURL
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
	allow := robots.TestAgent(targetURL, agent)
	return allow
}

func errCheck(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
