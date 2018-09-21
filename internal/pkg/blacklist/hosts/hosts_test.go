package hosts

import (
	"testing"
)

func TestSources(t *testing.T) {
	urls := Sources()
	if len(urls) < 2 {
		t.Errorf("expecting results but received none")
	}
}
