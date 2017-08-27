package onelike

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const hostUrl string = "http://onelike.tv"

const (
	mezzoUrl         string = hostUrl + "/mezzo.html"
	frameUrlPattern  string = "<iframe .+ name=\"frame\" src=\"(https?://.+)\" .+>"
	logoUrlPattern   string = "<img src=\"(/.+)\" border=\"0\" .+ width=\"70\" height=\"70\" .+>"
	streamUrlPattern string = "<param name=\"flashvars\" value=\".+file=(http?://.+)\">"
)

type channel struct {
	streamUrl string
	logoUrl   string
}

func Channel() string {
	mezzo := getChannel(mezzoUrl)

	return mezzo.streamUrl
}

func getChannel(url string) channel {
	bodyHtml := getBody(url, hostUrl)
	frameUrl := getUrl(bodyHtml, frameUrlPattern)
	logoUrl := hostUrl + getUrl(bodyHtml, logoUrlPattern)

	frameHtml := getBody(frameUrl, hostUrl)
	streamUrl := getUrl(frameHtml, streamUrlPattern)

	return channel{streamUrl, logoUrl}
}

func getBody(url, referer string) (bodyHtml string) {
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
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyHtml = string(bodyBytes)
	}

	return
}

func getUrl(html, pattern string) (matchedUrl string) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	matched := r.FindStringSubmatch(html)

	if len(matched) == 2 {
		matchedUrl = matched[1]
	}

	return
}
