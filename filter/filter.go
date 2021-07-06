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
		targetURL := getTargetURL(url)
		robotsRule := getRobotsRules(robotsURL)
		isAllowURL := isAllowURLWithRobots(robotsRule, agent, targetURL)
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
