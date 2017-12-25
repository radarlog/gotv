// Package onelike is a gotv's plugin implements parsing any channel on http://onelike.tv/
//
package onelike

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// the regex pattern for finding stream out
const streamUrlPattern = "file: '(https?://.+)'"

// find out channel's stream by parsing its page. All the pages are organized in the same way.
// A jQuery video player which contains a wanted stream is inserted to the page inside iframe tag,
// so fetching the stream is split into 2 steps: finding frame's URL and page's parsing by that URL
func FindStream(pageUrl string) string {
	frameUrl := getFrameUrl(pageUrl)
	streamUrl := getStreamUrl(frameUrl, pageUrl)

	return streamUrl
}

// find out an URL of the iframe
func getFrameUrl(pageUrl string) string {
	html, err := goquery.NewDocument(pageUrl)
	if err != nil {
		log.Fatal(err)
	}

	return html.Find("iframe[name=frame]").AttrOr("src", "")
}

// find out the stream inside iframe
func getStreamUrl(frameUrl string, referer string) (matchedUrl string) {
	response := request(frameUrl, referer)

	r := regexp.MustCompile(streamUrlPattern)
	matched := r.FindStringSubmatch(response)

	if len(matched) == 2 {
		matchedUrl = matched[1]
	}

	return
}

// iframe's URL is accessible only from channel's page, so it needs to be passed as a referer
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
