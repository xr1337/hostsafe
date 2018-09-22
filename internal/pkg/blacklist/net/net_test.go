package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	web := Web{}
	data, err := web.Download("http://www.google.com")
	assert.Nil(t, err)
	assert.True(t, len(data) > 1)
}
