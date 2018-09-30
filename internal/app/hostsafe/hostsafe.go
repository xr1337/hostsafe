package hostsafe

import (
	"fmt"
	"os"
	"strings"

	"github.com/xr1337/hostsafe/internal/pkg/blacklist/net"
)

// DownloadWorker function to download host and parse them. Removes commented lines
func DownloadWorker(url string, outChan chan string) {
	web := net.Web{}
	content, err := web.Download(url)
	if err != nil {
		fmt.Println("unable to download " + url)
		//panic(err)
		outChan <- ""
		return
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	goodLines := cleanHostEntry(lines)
	outChan <- strings.Join(goodLines, "\n")
}

func cleanHostEntry(entries []string) []string {
	cleanHosts := []string{}
	for _, line := range entries {
		l := strings.TrimSpace(line)
		if len(l) <= 0 {
			continue
		}
		if strings.HasPrefix(l, "#") {
			continue
		}
		if strings.HasPrefix(l, "127.0.0.1") {
			cleanHosts = append(cleanHosts, l)
			continue
		}
		if strings.HasPrefix(l, "0.0.0.0") {
			cleanHosts = append(cleanHosts, l)
			continue
		}
		if strings.HasPrefix(l, "0 ") {
			cleanHosts = append(cleanHosts, l)
			continue
		}
		if strings.HasPrefix(l, "::") {
			cleanHosts = append(cleanHosts, l)
			continue
		}
		cleanHosts = append(cleanHosts, "127.0.0.1 "+l)
	}
	return cleanHosts
}

// Process each host and adds them a temp file
func Process(inChan chan string, count int) string {
	filename := "/tmp/bad_hosts"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < count; i++ {
		text := <-inChan
		f.WriteString(text)
	}
	f.Sync()
	return filename
}

// ReplaceEtcHostFile /etc/host with the contents of filename
func ReplaceEtcHostFile(filename string) {
	if err := os.Rename(filename, "/etc/hosts"); err != nil {
		panic(err)
	}
}
