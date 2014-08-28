package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func DownloadURL(url string) string { //This function is taken from "Go-scrape"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents)
}

func deHTML(html string) string { //When/if I get this to work, i will make it another program. There is already a program called dehtml, but I haven't looked at how it works
	//I might use some regex here just for a quick solution and then use go.net/html later.
	re := regexp.MustCompile("<.*?>")
	return re.ReplaceAllString(html, "")
}

func main() {
	URL := "https://en.wikipedia.org/w/api.php?format=json&action=query&generator=random&prop=extracts&grnnamespace=0"
	wikipediaRaw := DownloadURL(URL)
	re := regexp.MustCompile(`(\"extract\"\:\")(.*?\.)`) //TODO: Parse this properly, without regular expressions. While it works for most of the cases I don't want to find one where it doesn't. And hey, it's JSON so yeah
	fmt.Printf("%s\n", deHTML(re.FindStringSubmatch(wikipediaRaw)[2]))
}
