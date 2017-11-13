package main

import (
    "net/http"
    "fmt"
    "io"
    "golang.org/x/net/html"
    "math/rand"
    "time"
    "strings"
)

//Take a URL, take all the external URLs on the page and insert it into the DB.
func Scrape(url string) {
    rand.Seed(time.Now().UnixNano())
    webId := GenRandID()
    links := Parse(Request(url))
    stmt := fmt.Sprintf("INSERT INTO webpages VALUES (%d, \"%s\")", webId, url)
    Execute(stmt)
    for _, element := range links {
        stmt = fmt.Sprintf("INSERT INTO weblinks VALUES (%d, \"%s\")", webId, element)
        Execute(stmt)
    }
}

//Generate a random ID between 1-1000 for the ID of a webpage.
func GenRandID() int {
    return rand.Intn(1000)
}

//Request a webpage.
func Request(url string) io.Reader {
    resp, _ := http.Get(url)
    return resp.Body
}

//Parse the body of the webpage and return the links throughout it.
func Parse(r io.Reader) []string {
    z := html.NewTokenizer(r)

    var links []string

    OUTERLOOP:
    for {
        iter := z.Next()
        switch {
            case iter == html.ErrorToken:
                break OUTERLOOP
            case iter == html.StartTagToken:
                t := z.Token()

                isAnchor := t.Data == "a"
                if isAnchor {
                    for _, a := range t.Attr {
                        if a.Key == "href" {
                            val := a.Val
                            if strings.Index(val, "http") == 0 {
                                links = append(links, a.Val)
                            }
                        }
                    }
                }
        }
    }

    return links
}
