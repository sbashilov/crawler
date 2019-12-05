package crawler

import (
	"errors"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Crawler custom errors
var (
	ErrEmptyURLHost = errors.New("provided url has empty host")
)

// Crawler represents a crawler that can be used to build hierarchical struct of links from the provided url
type Crawler struct {
	url          *url.URL
	visitedPages map[string]*Page
	mutex        *sync.Mutex
}

// New creates new crawler
func New(rawURL string) (*Crawler, error) {
	c := &Crawler{}
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if u.Host == "" {
		return nil, ErrEmptyURLHost
	}
	u.Fragment = ""
	c.url = u
	return c, nil
}

func (c *Crawler) init() {
	c.mutex = &sync.Mutex{}
	c.visitedPages = make(map[string]*Page)
}

// Crawl builds hierarchical struct of links from the provided url
// It returns pointer to Page struct
func (c *Crawler) Crawl() (*Page, error) {
	c.init()
	ch := make(chan *Page)
	r := Request{
		URL: c.url.String(),
	}
	go c.fetch(r, ch)
	page := <-ch
	return page, nil
}

func (c *Crawler) fetch(r Request, ch chan<- *Page) {
	c.mutex.Lock()
	if p, ok := c.visitedPages[trimURL(r.URL)]; ok {
		ch <- &Page{Link: p.Link}
		c.mutex.Unlock()
		return
	}
	p := &Page{Link: r.URL}
	c.visitedPages[trimURL(r.URL)] = p
	c.mutex.Unlock()
	d, err := r.Call()
	if err != nil {
		ch <- p
		return
	}
	uniqueCrossLinks := c.getUniqueCrossLinksFromPage(d, r)
	if len(uniqueCrossLinks) == 0 {
		ch <- p
		return
	}
	pCh := make(chan *Page, len(uniqueCrossLinks))
	p.CrossPages = make([]*Page, 0, len(uniqueCrossLinks))
	go func() {
		wg := sync.WaitGroup{}
		for l := range uniqueCrossLinks {
			wg.Add(1)
			r := Request{
				URL: l,
			}
			go func() {
				c.fetch(r, pCh)
				wg.Done()
			}()
		}
		wg.Wait()
		close(pCh)
	}()
	for cp := range pCh {
		p.CrossPages = append(p.CrossPages, cp)
	}
	ch <- p
	return

}

func (c *Crawler) getUniqueCrossLinksFromPage(d *goquery.Document, r Request) map[string]struct{} {
	uniqueCrossLinks := make(map[string]struct{})
	d.Find("a").Each(func(_ int, s *goquery.Selection) {
		var href string
		var ok bool
		if href, ok = s.Attr("href"); !ok {
			return
		}
		u, err := c.url.Parse(href)
		if err != nil {
			return
		}
		u.Fragment = ""
		if u.Host != c.url.Host {
			return
		}
		if trimURL(u.String()) == trimURL(r.URL) {
			return
		}
		if _, ok := uniqueCrossLinks[u.String()]; !ok {
			uniqueCrossLinks[u.String()] = struct{}{}
		}
	})
	return uniqueCrossLinks
}

func trimURL(url string) string {
	return strings.Trim(url, "/")
}
