package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Brief map[string]string

func (brief Brief) Time() time.Time {
	t, _ := time.ParseInLocation("20060102150405", brief["SEND_DATETIME"], time.Local)
	return t
}

func (brief Brief) Title() string {
	return brief["TITLE"]
}

func (brief Brief) URL() string {
	return brief["CONTENTS_LINK"]
}

func Briefs() ([]Brief, error) {
	resp, err := http.Get("http://www.yonhapnews.co.kr/data/sokbo_ticker.js")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	v := make(map[string][]Brief)
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}

	return v["DATA"], nil
}
