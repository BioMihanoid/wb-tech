package command

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Crawler struct {
	output   string
	maxDepth int
	client   *http.Client
	seen     sync.Map
	wg       sync.WaitGroup
	sem      chan struct{}
}

func NewCrawler(output string, depth, parallel, timeout int) *Crawler {
	return &Crawler{
		output:   output,
		maxDepth: depth,
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		sem: make(chan struct{}, parallel),
	}
}

func (c *Crawler) Crawl(startURL string) error {
	u, err := url.Parse(startURL)
	if err != nil {
		return err
	}
	host := strings.ReplaceAll(u.Host, ":", "_")
	if err := os.MkdirAll(filepath.Join(c.output, host), 0755); err != nil {
		return err
	}

	c.wg.Add(1)
	go c.fetch(startURL, 0)
	c.wg.Wait()
	return nil
}

func (c *Crawler) fetch(rawURL string, depth int) {
	defer c.wg.Done()

	if depth > c.maxDepth {
		return
	}
	if _, loaded := c.seen.LoadOrStore(rawURL, true); loaded {
		return
	}

	c.sem <- struct{}{}
	defer func() { <-c.sem }()

	resp, err := c.client.Get(rawURL)
	if err != nil {
		fmt.Println("error fetching", rawURL, err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status", rawURL, resp.Status)
		return
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}

	localPath := c.localPath(u)
	if err = os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
		fmt.Println("mkdir error", err)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "text/html") {
		modified := c.rewriteAndQueue(bodyBytes, u, depth)
		if err = os.WriteFile(localPath, modified, 0644); err != nil {
			fmt.Println("file write error", err)
		}
	} else {
		if err = os.WriteFile(localPath, bodyBytes, 0644); err != nil {
			fmt.Println("file write error", err)
		}
	}
}

// localPath формирует путь для сохранения URL
func (c *Crawler) localPath(u *url.URL) string {
	host := strings.ReplaceAll(u.Host, ":", "_")

	path := strings.TrimPrefix(u.Path, "/")
	if path == "" {
		path = "index.html"
	}

	localPath := filepath.Join(c.output, host, path)

	if strings.HasSuffix(u.Path, "/") || filepath.Ext(path) == "" {
		localPath = filepath.Join(c.output, host, path, "index.html")
	}
	return localPath
}

// rewriteAndQueue парсит HTML, переписывает ссылки и ставит новые задачи в очередь
func (c *Crawler) rewriteAndQueue(data []byte, base *url.URL, depth int) []byte {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return data
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attribute string
			switch n.Data {
			case "a":
				attribute = "href"
			case "img", "script":
				attribute = "src"
			case "link":
				attribute = "href"
			}
			if attribute != "" {
				for i := range n.Attr {
					if n.Attr[i].Key == attribute {
						link := n.Attr[i].Val
						abs := c.resolveURL(base, link)
						if abs == "" {
							continue
						}
						if strings.HasPrefix(abs, base.Scheme+"://"+base.Host) {
							parsed, _ := url.Parse(abs)
							local := c.localPath(parsed)
							rel, _ := filepath.Rel(filepath.Dir(c.localPath(base)), local)
							n.Attr[i].Val = rel

							c.wg.Add(1)
							go c.fetch(abs, depth+1)
						}
					}
				}
			}
		}
		for ch := n.FirstChild; ch != nil; ch = ch.NextSibling {
			f(ch)
		}
	}
	f(doc)

	var buf bytes.Buffer
	if err = html.Render(&buf, doc); err != nil {
		return data
	}
	return buf.Bytes()
}

func (c *Crawler) resolveURL(base *url.URL, href string) string {
	u, err := url.Parse(href)
	if err != nil {
		return ""
	}
	return base.ResolveReference(u).String()
}
