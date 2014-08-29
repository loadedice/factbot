package main

import (
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var verbose = flag.Bool("v", false, "Verbose logging when errors happen")

func verboseLog(verb bool, message string) {
	if verb {
		log.Println(message)
	}
	return
}

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

func deHTML(html string) string {
	//When/if I get this to work, i will make it another program. There is already a program called dehtml, but I haven't looked at how it works
	//I might use some regex here just for a quick solution and then use go.net/html later.
	re := regexp.MustCompile("<.*?>")
	return re.ReplaceAllString(html, "")
}

func main() {
	flag.Parse()
	verboseLog(*verbose, "Verbose logging enabled.")
	URL := "https://en.wikipedia.org/w/api.php?format=json&action=query&generator=random&prop=extracts&grnnamespace=0"
	wikipediaRaw := DownloadURL(URL)
	re := regexp.MustCompile(`(\"extract\"\:\")(.*?\.)`)
	//TODO: Parse this properly, without regular expressions. While it works for most of the cases I don't want to find one where it doesn't. And hey, it's JSON so yeah
	fact := html.UnescapeString(deHTML(re.FindStringSubmatch(wikipediaRaw)[2]))
	for {
		if len([]rune(fact)) > 10 {
			fmt.Printf("%s\n", fact)
			break //yea, I need to use the proper syntax rather than this cheap hack
		} else {
			verboseLog(*verbose, "Fact was shorter than 10 runes. Will attempt to get another fact")
			//it actually doesn't... Yet... so we'll break
			break
		}
	}
}
