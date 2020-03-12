# Web scrapper #

## Introduction ##

This scrapes Item names and prices from a given website. In this case, we scrape Graphics Card prices from www.newegg.com.

## Behind the scenes ##

* The web scrapper is written in go using golang.org/x/net/html package.

* Uses the html package to parse the website HTML body. 

* Once the HTML is parsed, and the structure of the website is known, the div that contains the name/price can be easily traversed to. 

## What did I Learn ##

* Web scraping basics.

* Learned about the html package.
