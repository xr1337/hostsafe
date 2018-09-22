package hosts

import (
	"strings"
)

// Resource is a interface that exposes methods to retrieve resources
type Resource interface {
	Download(url string) (text string, err error)
}

func downloadsources(r Resource) []string {
	firebogurl := "https://v.firebog.net/hosts/lists.php?type=tick"
	text, err := r.Download(firebogurl)
	if err != nil {
		panic(err)
	}
	urls := strings.Split(strings.TrimSpace(text), "\n")

	return urls
}

// Sources return a list of urls that has hosts
func Sources(r Resource) []string {
	urls := downloadsources(r)
	sourceURLs := urls[:]
	sourceURLs = append(sourceURLs, "https://someonewhocares.org/hosts/hosts")
	return sourceURLs
}
