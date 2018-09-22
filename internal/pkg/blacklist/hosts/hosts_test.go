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

func TestDownload(t *testing.T) {
	m := MockResource{}
	urls := downloadsources(m)
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
	downloadsources(m)
}

// MockResource
func (MockResource) Download(url string) (text string, err error) {
	return "this\nis\ncool\n", nil
}

// BadMockResource
func (BadMockResource) Download(url string) (text string, err error) {
	return "", errors.New("mock error")
}
