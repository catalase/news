package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func news(name, text string, when int64, url string) {
	cmd := exec.Command("python", "news.py",
		name, text, strconv.FormatInt(when, 10), url)
	cmd.Start()
}

func hasNew(old map[string]bool, briefs []Brief) bool {
	for _, brief := range briefs {
		if !old[brief.URL()] {
			return true
		}
	}

	return false
}

func makeOld(briefs []Brief) map[string]bool {
	old := make(map[string]bool)
	for _, brief := range briefs {
		old[brief.URL()] = true
	}
	return old
}

func main() {
	var old map[string]bool
	tick := time.Tick(5 * time.Second)

Retry:
	if briefs, err := Briefs(); err != nil {
		<-tick
		goto Retry
	} else {
		old = makeOld(briefs)
	}

	for range tick {
		briefs, err := Briefs()
		if err != nil {
			continue
		}

		if !hasNew(old, briefs) {
			continue
		}

		for _, brief := range briefs {
			if url := brief.URL(); !old[url] {
				fmt.Println(time.Now(), brief.Title())
				news(
					brief.Title(),
					brief.Title(),
					brief.Time().Unix(),
					url,
				)
			}
		}

		old = makeOld(briefs)
	}
}
