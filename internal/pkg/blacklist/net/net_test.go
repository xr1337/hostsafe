package net

import "testing"

func TestDownload(t *testing.T) {
	web := Web{}
	data, err := web.Download("http://www.google.com")
	if err != nil {
		t.Errorf("download error %s", err)
	}
	if len(data) < 1 {
		t.Error("download empty")
	}
}
