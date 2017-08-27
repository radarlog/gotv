package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

const (
    mezzoRequestUrl  string = "http://suppor6k.bget.ru/onelike/mezzo.php"
    mezzoRefererUrl  string = "http://onelike.tv/mezzo.html"
    streamUrlPattern string = "http://s2a.privit.pro:8081/mezzo/index.m3u8\\?wmsAuthSign=(.+==)"
)

func main() {
    body := getBody(mezzoRequestUrl, mezzoRefererUrl)

    streamUrl := getLink(body)

    fmt.Println(streamUrl)
}

func getBody(url string, referer string) string {
    request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("Referer", referer)

    client := http.Client{}
    resp, _ := client.Do(request)
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        return string(bodyBytes)
    }

    return ""
}

func getLink(body string) string {
    r, _ := regexp.Compile(streamUrlPattern)
    streamUrl := r.FindString(body)

    return streamUrl
}
