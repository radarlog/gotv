package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	channelUrl       string = "http://onelike.tv/mezzo.html"
	frameUrlPattern  string = "<iframe .+ name=\"frame\" src=\"(https?://.+)\" .+>"
	streamUrlPattern string = "<param name=\"flashvars\" value=\".+file=(http?://.+)\">"
)

func main() {
	bodyHtml := getBody(channelUrl, channelUrl)
	frameUrl := getUrl(bodyHtml, frameUrlPattern)

	bodyHtml = getBody(frameUrl, channelUrl)
	streamUrl := getUrl(bodyHtml, streamUrlPattern)

	fmt.Println(streamUrl)
}

func getBody(url, referer string) (bodyHtml string) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Referer", referer)

	client := http.Client{}
	resp, _ := client.Do(request)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyHtml = string(bodyBytes)
	}

	return
}

func getUrl(body, pattern string) (matchedUrl string) {
	r, _ := regexp.Compile(pattern)
	matched := r.FindStringSubmatch(body)

	if len(matched) == 2 {
		matchedUrl = matched[1]
	}

	return
}
