package onelike

import (
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Tv struct {
	Mezzo      Channel `yaml:"mezzo"`
	NatGeoWild Channel `yaml:"nat-geo-wild"`
}

type Channel struct {
	Name      string `yaml:"name"`
	PageUrl   string `yaml:"page_url"`
	LogoUrl   string `yaml:"logo_url"`
	StreamUrl string
}

func (c Channel) Stream() string {
	frameUrl := getFrameUrl(c.PageUrl, c.PageUrl)
	streamUrl := getStreamUrl(frameUrl, c.PageUrl)

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
