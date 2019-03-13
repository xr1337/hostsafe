package hostsafe

import (
	"bufio"
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
		outChan <- ""
		return
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	goodLines := cleanHostEntry(lines)
	outChan <- strings.Join(goodLines, "\n")
}

func extractHost(entry string) string {
	if len(entry) <= 0 {
		return ""
	}
	if strings.Contains(entry, "#") {
		return ""
	}
	entry = strings.Replace(entry, "127.0.0.1", "", -1)
	parts := strings.Split(strings.TrimSpace(entry), " ")
	if len(parts) == 1 {
		return parts[0]
	}
	return parts[1]
}

func cleanHostEntry(entries []string) []string {
	cleanHosts := []string{}
	for _, line := range entries {
		l := strings.TrimSpace(line)
		host := extractHost(l)
		if len(host) > 0 {
			cleanHosts = append(cleanHosts, host)
		}
	}
	return cleanHosts
}

// Process each host and adds them a temp file
func Process(inChan chan string, count int, exclude map[string]string) string {
	var m map[string]string
	m = make(map[string]string)
	for i := 0; i < count; i++ {
		text := <-inChan
		scanner := bufio.NewScanner(strings.NewReader(text))
		for scanner.Scan() {
			host := scanner.Text()
			if _, ok := exclude[host]; !ok {
				m[host] = ""
			}
		}
	}

	filename := "/tmp/bad_hosts"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for key := range m {
		f.WriteString("127.0.0.1 " + key + "\n")
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
