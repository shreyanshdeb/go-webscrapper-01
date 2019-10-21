package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type item struct {
	name  string
	price string
}

func main() {
	for _, v := range getURLs() {
		fmt.Println(dataExtrator(fetchHTML(v)))
	}
}

func getURLs() []string {
	urls := []string{"https://www.newegg.com/Video-Cards-Video-Devices/Category/ID-38?Tpk=graphic%20cards"}
	return urls
}

func fetchHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "Error getting response"
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading body"
	}
	fmt.Println("HTML Fetched")
	return string(html)
}

func dataExtrator(htmlData string) *[]item {
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		fmt.Println("Error parsing html", err)
	}
	fmt.Println("HTML Parsed")
	itemFound := new([]item)
	var f func(*html.Node, *[]item)
	itemName := ""
	f = func(n *html.Node, itemFound *[]item) {
		itemPrice := ""
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, v := range n.Attr {
				attr := html.Attribute{
					Key: "class",
					Val: "item-title",
				}
				if v == attr {
					itemName = n.FirstChild.NextSibling.Data
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "strong" && n.Parent.Data == "li" {
			for _, v := range n.Parent.Attr {
				attr := html.Attribute{
					Key: "class",
					Val: "price-current",
				}
				if v == attr {
					itemPrice = n.FirstChild.Data
					itemNew := item{
						name:  itemName,
						price: itemPrice,
					}
					itemName = ""
					*itemFound = append(*itemFound, itemNew)
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, itemFound)
		}
	}
	f(doc, itemFound)

	return itemFound
}
