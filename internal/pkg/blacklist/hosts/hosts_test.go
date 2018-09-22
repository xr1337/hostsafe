package hosts

import (
	"errors"
	"testing"
)

type MockResource struct {
}

type BadMockResource struct {
}

func TestSources(t *testing.T) {
	m := MockResource{}
	urls := Sources(m)
	if len(urls) < 2 {
		t.Errorf("expecting results but received none")
	}
}

func TestValidateURL(t *testing.T) {
	var table = []struct {
		urls     []string
		expected int
	}{
		{[]string{"www.google.com", "http://badhost.yahoo.com"}, 1},
		{[]string{"https://www.google.com", "http://www.booking.com"}, 2},
		{[]string{"https://www.booking.com"}, 1},
		{[]string{"a", "b"}, 0},
	}
	for _, s := range table {
		if actual := filterValidUrls(s.urls); len(actual) != s.expected {
			t.Errorf("expected %s to have %d items", s.urls, s.expected)
		}
	}
}

func TestDownload(t *testing.T) {
	m := MockResource{}
	urls := downloadSources(m)
	if len(urls) < 1 {
		t.Errorf("expected urls to be more than 1")
	}
}

func TestPanicDownload(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m := BadMockResource{}
	downloadSources(m)
}

// MockResource
func (MockResource) Download(url string) (text string, err error) {
	return "https://www.google.com\nhttps://www.booking.com\n", nil
}

// BadMockResource
func (BadMockResource) Download(url string) (text string, err error) {
	return "", errors.New("mock error")
}
