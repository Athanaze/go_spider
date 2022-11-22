package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/corpix/uarand"
	"github.com/gocolly/colly/v2"
	"golang.org/x/net/html"
)

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func completeURL(domain, u string) string {
	defer func() {
		if err := recover(); err != nil {
			//log.Println("panic occurred:", err)
		}
	}()
	domain = TrimSuffix(domain, "/")
	if u[0:1] == "/" {
		return domain + u
	} else {
		return u
	}
}
func repetition(st string) map[string]int {
	input := strings.Split(st, "/")
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

// Collect all links from response body and return it as an array of strings
func getLinks(u string) []string {
	resp, err := http.Get(u)
	if err != nil {
		return []string{""}
	}
	body := resp.Body
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			//log.Println(links)
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, completeURL(u, attr.Val) /*, attr.Val*/)
					}

				}
			}

		}
	}
}
func v(u string, parentURL string) {
	if u != parentURL && !tooManyTimes(u) {
		c := colly.NewCollector(colly.IgnoreRobotsTxt(), colly.DisallowedDomains(
			"en.wikipedia.org"), colly.DisallowedDomains(
			"fr.wikipedia.org"), colly.DisallowedDomains(
			"de.wikipedia.org"), colly.DisallowedDomains(
			"ja.wikipedia.org"), colly.DisallowedDomains(
			"es.wikipedia.org"), colly.DisallowedDomains(
			"ru.wikipedia.org"), colly.DisallowedDomains(
			"pt.wikipedia.org"), colly.DisallowedDomains(
			"m.facebook.com"))

		// Find and visit all links
		c.OnHTML("a", func(e *colly.HTMLElement) {
			go func() {
				c.UserAgent = uarand.GetRandom()
				a := completeURL(u, e.Attr("href"))
				fmt.Println(a)
				//fmt.Println("Parent :", u, " child : ", a)
				v(a, u)
			}()

		})

		c.Visit(u)
	}

}
func tooManyTimes(s string) bool {
	r := false
	for _, element := range repetition(s) {
		if element > 3 {
			r = true
		}
	}
	return r
}
func main() {
	U2 := []string{"https://www.coindesk.com/", "http://www.reuters.com/world/", "https://nitter.net/chepurnoy?lang=en", "https://teddit.net/r/cryptocurrency", "https://stackoverflow.com/questions/44039223/golang-why-are-goroutines-not-running-in-parallel", "https://github.com/corpix/uarand", "https://www.reply.com/alpha-reply/en/content/go-concurrency-with-mutex", "https://www.digitalocean.com/community/tutorials/handling-panics-in-go", "https://vorozhko.net/get-all-links-from-html-page-with-go-lang", "https://ngrok.com/", "https://teddit.net/r/worldnews", "https://teddit.net/r/news", "https://www.coinbureau.com/newsletters/the-fed-is-on-a-mission-of-pain/"}

	//go v(U2[0], "a")
	//go v(U2[0], "a")
	go v(U2[2], "a")
	for {

	}

}
