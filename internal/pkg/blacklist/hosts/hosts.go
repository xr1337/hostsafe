package hosts

import (
	"net/url"
	"strings"
)

// Resource is a interface that exposes methods to retrieve resources
type Resource interface {
	Download(url string) (text string, err error)
}

// Sources return a list of urls that has hosts
func Sources(r Resource) []string {
	urls := downloadSources(r)
	sourceURLs := append(urls[:], "https://someonewhocares.org/hosts/hosts")
	sourceURLs = filterValidUrls(sourceURLs)
	return sourceURLs
}

func downloadSources(r Resource) []string {
	firebogurl := "https://v.firebog.net/hosts/lists.php?type=tick"
	text, err := r.Download(firebogurl)
	if err != nil {
		panic(err)
	}
	urls := strings.Split(strings.TrimSpace(text), "\n")
	return urls
}

func filterValidUrls(urls []string) []string {
	cleanUrls := []string{}
	for _, urlString := range urls {
		if _, err := url.ParseRequestURI(urlString); err == nil {
			cleanUrls = append(cleanUrls, urlString)
		}
	}
	return cleanUrls
}
