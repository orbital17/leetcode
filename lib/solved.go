package lib

import (
	"io/ioutil"
)

func calcSolved() int {
	files, _ := ioutil.ReadDir("..")
	dict := make(map[string]bool)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		k := file.Name()[:3]
		dict[k] = true
	}
	return len(dict)
}
