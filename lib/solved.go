package lib

import (
	"fmt"
	"io/ioutil"
	"os"
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

func Refactor() error {
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		err := moveToFolder(file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func moveToFolder(filename string) error {
	f := filename[:len(filename)-3]
	if f[len(f)-5:] == "_test" {
		f = f[:len(f)-5]
	}
	err := os.MkdirAll(f, 0755)
	if err != nil {
		return err
	}
	err = os.Rename(filename, fmt.Sprintf("%s/%s", f, filename))
	return err
}
