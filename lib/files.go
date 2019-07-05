package lib

import (
	"fmt"
	"io/ioutil"
	"os"
)

func NewProblem(name string) (err error) {
	err = os.MkdirAll(name, 0755)
	if err != nil {
		return
	}
	f, err := os.Create(fmt.Sprintf("%s/%s.go", name, name))
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.WriteString("package leetcode\n")
	if err != nil {
		return
	}
	return
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
