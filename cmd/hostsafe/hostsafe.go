package main

import (
	"github.com/xr1337/hostsafe/internal/app/hostsafe"
	"github.com/xr1337/hostsafe/internal/pkg/blacklist/hosts"
)

func main() {
	output := make(chan string, 5)
	jobs := hosts.Sources()

	for _, url := range jobs {
		go hostsafe.DownloadWorker(url, output)
	}
	filename := hostsafe.Process(output, len(jobs))
	hostsafe.ReplaceEtcHostFile(filename)
}
