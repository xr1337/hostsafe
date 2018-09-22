package hosts

import (
	"testing"
)

type MockResource struct {
}

func TestSources(t *testing.T) {
	m := MockResource{}
	urls := Sources(m)

	if len(urls) < 2 {
		t.Errorf("expecting results but received none")
	}
}

func TestDownload(t *testing.T) {
	m := MockResource{}
	urls := downloadsources(m)
	if len(urls) < 1 {
		t.Errorf("expected urls to be more than 1")
	}
}

// mock items
func (MockResource) Download(url string) (text string, err error) {
	return "this\nis\ncool\n", nil
}
