package onelike

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

const streamUrlPattern = "file: '(https?://.+)'"

func FindStream(pageUrl string) string {
	frameUrl := getFrameUrl(pageUrl, pageUrl)
	streamUrl := getStreamUrl(frameUrl, pageUrl)

	return streamUrl
}

func getFrameUrl(pageUrl string, referer string) string {
	html, err := goquery.NewDocument(pageUrl)
	if err != nil {
		log.Fatal(err)
	}

	return html.Find("iframe[name=frame]").AttrOr("src", "")
}

func getStreamUrl(pageUrl string, referer string) (matchedUrl string) {
	response := request(pageUrl, referer)

	r := regexp.MustCompile(streamUrlPattern)
	matched := r.FindStringSubmatch(response)

	if len(matched) == 2 {
		matchedUrl = matched[1]
	}

	return
}

func request(url, referer string) (html string) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Referer", referer)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBite, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		html = string(bodyBite)
	}

	return html
}
