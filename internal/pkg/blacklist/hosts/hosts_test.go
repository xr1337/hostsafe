package hosts

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockResource struct {
}

type BadMockResource struct {
}

func TestSources(t *testing.T) {
	m := MockResource{}
	urls := Sources(m)
	assert.True(t, len(urls) > 2)
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
		assert.Equal(t, s.expected, len(filterValidUrls(s.urls)))
	}
}

func TestDownload(t *testing.T) {
	m := MockResource{}
	urls := downloadSources(m)
	assert.True(t, len(urls) > 1)
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
