package net

import (
	"io/ioutil"
	"net/http"
)

func Download(url string) (text string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return text, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return text, err
	}
	text = string(contents)
	return text, nil
}
