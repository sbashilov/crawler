package main

import (
	"flag"
	"log"
	"os"

	"github.com/sbashilov/crawler"
)

func main() {
	out := os.Stdout
	url := flag.String("url", "", "Crawler will start with this url")
	depth := flag.Int("depth", 10, "Page struct print depth")
	flag.Parse()
	c, err := crawler.New(*url)
	if err != nil {
		log.Fatal(err)
	}
	p, err := c.Crawl()
	if err != nil {
		log.Fatal(err)
	}
	p.Print(out, *depth)
}
