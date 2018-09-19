package hosts

import (
	"github.com/xr1337/hostsafe/internal/pkg/blacklist/net"
	"strings"
)

func downloadsources() []string {
	firebogurl := "https://v.firebog.net/hosts/lists.php?type=tick"
	text, err := net.Download(firebogurl)
	if err != nil {
		panic(err)
	}
	urls := strings.Split(strings.TrimSpace(text), "\n")

	return urls
}

func Sources() []string {
	urls := downloadsources()
	sourceURLs := urls[:]
	sourceURLs = append(sourceURLs, "https://someonewhocares.org/hosts/hosts")
	return sourceURLs
}
