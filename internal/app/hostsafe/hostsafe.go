package hostsafe

import (
	"github.com/xr1337/hostsafe/internal/pkg/blacklist/net"
	"os"
	"strings"
)

func DownloadWorker(url string, outChan chan string) {
	content, err := net.Download(url)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	goodLines := []string{}
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if len(l) <= 0 {
			continue
		}
		if strings.HasPrefix(l, "#") {
			continue
		}
		if strings.HasPrefix(l, "127.0.0.1") {
			goodLines = append(goodLines, l)
			continue
		}
		if strings.HasPrefix(l, "0.0.0.0") {
			goodLines = append(goodLines, l)
			continue
		}
		if strings.HasPrefix(l, "0 ") {
			goodLines = append(goodLines, l)
			continue
		}
		if strings.HasPrefix(l, "::") {
			goodLines = append(goodLines, l)
			continue
		}
		goodLines = append(goodLines, "127.0.0.1 "+l)
	}
	outChan <- strings.Join(goodLines, "\n")
}

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

func ReplaceEtcHostFile(filename string) {
	if err := os.Rename(filename, "/etc/hosts"); err != nil {
		panic(err)
	}
}
