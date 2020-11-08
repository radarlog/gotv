// Package onlytv is a gotv plugin parsing channels on http://only-tv.org/
//
package onlytv

import (
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// the regex pattern for finding stream out
const streamUrlPattern = "file=(https?://.+?)\""

// find out channel's stream by parsing its page. All the pages are organized in the same way.
// A jQuery video player which contains a wanted stream is inserted to the page inside iframe tag,
// so fetching the stream is split into 2 steps: finding frame's URL and page's parsing by that URL
func FindStream(pageUrl string) string {
	frameUrl := getFrameUrl(pageUrl)
	streamUrl := getStreamUrl(frameUrl)

	return streamUrl
}

// find out an URL of the iframe
func getFrameUrl(pageUrl string) string {
	html := request(pageUrl)

	return html.Find("iframe").AttrOr("src", "")
}

// find out the stream inside iframe
func getStreamUrl(frameUrl string) (matchedUrl string) {
	html := request(frameUrl)

	player, err := html.Find("div.player").Html()
	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(streamUrlPattern)
	matched := r.FindStringSubmatch(player)

	if len(matched) == 2 {
		matchedUrl = matched[1]
	}

	return
}

// URL is accessible only from channel's page, so it needs to be passed as a referer
func request(url string) (doc *goquery.Document) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Referer", url)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		d, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc = d
	}

	return
}
