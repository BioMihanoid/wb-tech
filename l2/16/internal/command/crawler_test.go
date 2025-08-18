package command

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func tempDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "wget_test")
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = os.RemoveAll(dir)
	})
	return dir
}

func TestCrawler_FetchSingleHTML(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><body><h1>Hello</h1></body></html>`))
	}))
	defer ts.Close()

	dir := tempDir(t)
	c := NewCrawler(dir, 1, 2, 5)

	err := c.Crawl(ts.URL)
	require.NoError(t, err)

	u, _ := url.Parse(ts.URL)
	local := c.localPath(u)

	data, err := os.ReadFile(local)
	require.NoError(t, err)
	require.Contains(t, string(data), "Hello")
}

func TestCrawler_RewriteLinksAndDownloadResources(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(`
				<html>
					<body>
						<a href="/page2">Go next</a>
						<img src="/img.png">
					</body>
				</html>`))
		case "/page2":
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(`<html><body><p>Second page</p></body></html>`))
		case "/img.png":
			w.Header().Set("Content-Type", "image/png")
			_, _ = w.Write([]byte("PNGDATA"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	dir := tempDir(t)
	c := NewCrawler(dir, 2, 2, 5)

	err := c.Crawl(ts.URL)
	require.NoError(t, err)

	u, _ := url.Parse(ts.URL)
	index := c.localPath(u)
	data, err := os.ReadFile(index)
	require.NoError(t, err)
	htmlStr := string(data)

	require.NotContains(t, htmlStr, `href="/page2"`)
	require.Contains(t, htmlStr, "index.html")

	u2, _ := url.Parse(ts.URL + "/page2")
	require.FileExists(t, c.localPath(u2))

	u3, _ := url.Parse(ts.URL + "/img.png")
	require.FileExists(t, c.localPath(u3))
}

func TestCrawler_MaxDepth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(`<a href="/page2">p2</a>`))
		case "/page2":
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(`<a href="/page3">p3</a>`))
		case "/page3":
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(`<p>deep page</p>`))
		}
	}))
	defer ts.Close()

	dir := tempDir(t)
	c := NewCrawler(dir, 1, 2, 5)

	err := c.Crawl(ts.URL)
	require.NoError(t, err)

	u, _ := url.Parse(ts.URL)
	require.FileExists(t, c.localPath(u))

	u2, _ := url.Parse(ts.URL + "/page2")
	require.FileExists(t, c.localPath(u2))

	u3, _ := url.Parse(ts.URL + "/page3")
	_, err = os.Stat(c.localPath(u3))
	require.Error(t, err)
	require.True(t, os.IsNotExist(err))
}

func TestCrawler_NonHTMLResource(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("hello text"))
	}))
	defer ts.Close()

	dir := tempDir(t)
	c := NewCrawler(dir, 1, 2, 5)

	err := c.Crawl(ts.URL)
	require.NoError(t, err)

	u, _ := url.Parse(ts.URL)
	data, err := os.ReadFile(c.localPath(u))
	require.NoError(t, err)
	require.Equal(t, "hello text", strings.TrimSpace(string(data)))
}
