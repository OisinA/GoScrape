package main

import (
    "log"
    "net/http"
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    toScrape := receiveInput()
    log.Println("Setting up database.")
    Setup()
    log.Println("Setting up web server.")

    for _, element := range toScrape {
        Scrape(element)
    }

    http.HandleFunc("/", ServePage)
    http.ListenAndServe(":300", nil)
}

//Take in the input from the user. It returns a list of links.
func receiveInput() []string {
    fmt.Println("Input each URL to be scraped (press ENTER to stop input): ")
    reader := bufio.NewReader(os.Stdin)
    links := make([]string, 0)
    terminate := false
    for {
        if terminate {
            break
        }

        text, _ := reader.ReadString('\n')
        if text == "\r\n" {
            terminate = false
            break
        }
        if strings.Index(text, "http") == 0 {
            links = append(links, text[:len(text)-2])
        } else {
            fmt.Printf("Incorrect syntax.")
        }
    }

    return links
}

//Prepare and serve the page to the user.
func ServePage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    rows, err := Database().Query("SELECT id, url FROM webpages")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Fprintf(w, "<ul>")
    for rows.Next() {
        var display string = "<li>%d</li><li><a href=\"%s\">%s</a>"
        var id int
        var webpage string
        err = rows.Scan(&id, &webpage)
        if err != nil {
            log.Fatal(err)
        }
        display = fmt.Sprintf(display, id, webpage, webpage)
        fmt.Fprintf(w, display)
        fmt.Fprintf(w, "<ul>")

        stmt := fmt.Sprintf("SELECT id, url FROM webLinks WHERE id=%d", id)
        webLinks, err := Database().Query(stmt)
        if err != nil {
            log.Fatal(err)
        }

        for webLinks.Next() {
            var link string
            var id int
            err = webLinks.Scan(&id, &link)
            if err != nil {
                log.Fatal(err)
            }
            display = "<li><a href=\"%s\">%s</a></li>"
            display = fmt.Sprintf(display, link, link)

            fmt.Fprintf(w, display)
        }

        fmt.Fprintf(w, "</ul></li>")
    }
    fmt.Fprintf(w, "</ul>")
}
