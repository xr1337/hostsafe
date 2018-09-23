package hostsafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostValidate(t *testing.T) {
	var table = []struct {
		items    []string
		expected int
	}{
		{[]string{"127.0.0.1 badhhost.com"}, 1},
		{[]string{"#this is a comment"}, 0},
		{[]string{":: abc.com"}, 1},
		{[]string{"0 abc.com"}, 1},
		{[]string{"0.0.0.0 badhhost.com"}, 1},
		{[]string{"badhhost.com"}, 1},
	}
	for _, row := range table {
		assert.Equal(t, len(cleanHostEntry(row.items)), row.expected)
	}

}
