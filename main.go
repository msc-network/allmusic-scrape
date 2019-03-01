package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/puerkitobio/goquery"
	flag "github.com/spf13/pflag"
)

var (
	id       string
	url      string
	parent   string
	page     string
	filename string
)

func main() {
	fmt.Println("Welcome to the allmusic Album page scraper")

	flag.Parse()

	if flag.NFlag() == 0 {
		printHelp()
	}

	if id == "" {
		fmt.Printf("Album ID is missing, please re-run\n")
		os.Exit(1)
	}
	err := ioutil.WriteFile(filename, nil, 0754)
	check(err)
	scrapeAllMusic(id)
}

func scrapeAllMusic(id string) {
	url = "https://www.allmusic.com/album/"
	fmt.Println("Scraping album with the ID: ", id)
	page := url + id

	parseDoc(string(page))
}

func init() {
	flag.StringVarP(&id, "id", "i", "", "Pass the ID of the allmusic album page")
	flag.StringVarP(&filename, "filename", "f", "data.txt", "Pass in an optional filename")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printHelp() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func parseDoc(page string) {
	doc, err := goquery.NewDocument(page)

	artist := strings.Trim(doc.Find(".album-artist span a").Text(), "  ")
	album := strings.Trim(doc.Find(".album-title").Text(), " ")
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	f.WriteString(string(artist + " - " + album + "\n"))

	// write title and composer
	doc.Find(".title-composer").Each(func(index int, item *goquery.Selection) {
		title := strings.Trim(item.Find(".title a").Text(), " ")
		composer := strings.Trim(item.Find(".composer a").Text(), " ")
		writeRelease(artist, title, composer)
	})
	check(err)
}

func writeRelease(a, t, c string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	f.WriteString(string(a + " - " + t + " [" + c + "]" + "\n"))

	check(err)
	return
}
