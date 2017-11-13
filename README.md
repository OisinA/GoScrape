# GoScrape
I wrote this basic scraping tool and web display in order to practice the basics of Go. It uses the net/http package to display the results of the scrape. It uses SQLite to store the results of the scrape. At present, it just searches the webpage for external URLs and displays them in a list.

## To use
You have to build the files using a Go compiler. Once built, run the program and input each URL to be scraped. It will then run through each one and the results can then be found through a web browser by going to localhost/:300.
