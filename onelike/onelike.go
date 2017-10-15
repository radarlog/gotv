package onelike

import (
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func FindStream(pageUrl string) string {
	frameUrl := getFrameUrl(pageUrl, pageUrl)
	streamUrl := getStreamUrl(frameUrl, pageUrl)

	return streamUrl
}

func getFrameUrl(pageUrl string, referer string) string {
	html := request(pageUrl, referer)

	return html.Find("iframe[name=frame]").AttrOr("src", "")
}

func getStreamUrl(pageUrl string, referer string) string {
	html := request(pageUrl, referer)
	param := html.Find("object param[name=flashvars]").AttrOr("value", "")

	values, err := url.ParseQuery(param)
	if err != nil {
		log.Fatal(err)
	}

	return values.Get("file")
}

func request(url, referer string) (html *goquery.Document) {
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

	if resp.StatusCode == 200 {
		html, err = goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}

	return html
}
