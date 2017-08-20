package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
)

const (
    mezzoRequestUrl string = "http://suppor6k.bget.ru/onelike/mezzo.php"
    mezzoRefererUrl string = "http://onelike.tv/mezzo.html"
)

func main() {
    html := getBody(mezzoRequestUrl, mezzoRefererUrl)

    fmt.Print(string(html))
}

func getBody(url string, referer string) []byte {
    client := http.Client{}

    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        panic(err)
    }

    request.Header.Set("Referer", referer)

    response, err := client.Do(request)
    if err != nil {
        panic(err)
    }

    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    return body
}
