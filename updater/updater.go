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
	body := getBody(channelUrl, channelUrl)
	frameUrl := getLink(body, frameUrlPattern)

	body = getBody(frameUrl, channelUrl)
	streamUrl := getLink(body, streamUrlPattern)

	fmt.Println(streamUrl)
}

func getBody(url string, referer string) string {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Referer", referer)

	client := http.Client{}
	resp, _ := client.Do(request)
	defer resp.Body.Close()

	bodyString := ""
	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}

	return bodyString
}

func getLink(body string, pattern string) string {
	r, _ := regexp.Compile(pattern)
	matched := r.FindStringSubmatch(body)

	streamUrl := ""
	if len(matched) == 2 {
		streamUrl = matched[1]
	}

	return streamUrl
}
