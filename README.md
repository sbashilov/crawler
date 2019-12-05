#Crawler

###Usage

`c, err := crawler.New("https://gobyexample.com")`

`p, err := c.Crawl()`

`p.Print(out, *depth)`


###For tests you can run

`go run cmd/main.go -url=https://gobyexample.com -depth=4`