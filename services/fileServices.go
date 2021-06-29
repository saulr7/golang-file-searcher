package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	m "searcher/models"
	"strings"
	"sync"
)

var (
	matches   []string
	waitgroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func Search(dir, fileName string) {

	waitgroup.Add(1)
	go FileSearch(dir, fileName)
	waitgroup.Wait()

}

func ShowMatches() {
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}

func FileSearch(root, fileName string) {

	fmt.Println("Searching...")

	files, _ := ioutil.ReadDir(root)

	for _, file := range files {

		fn := file.Name()

		if strings.Contains(strings.ToLower(fn), strings.ToLower(fileName)) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, fn))
			lock.Unlock()
		}

		if file.IsDir() {
			waitgroup.Add(1)
			go FileSearch(filepath.Join(root, fn), fileName)
		}

	}

	waitgroup.Done()

}

func ReadInputs() (m.InputModel, error) {

	args := os.Args[1:]

	var input m.InputModel

	if len(args) < 2 {
		return input, errors.New("expected dir and filename")
	}
	input.Dir, input.FileName = args[0], args[1]

	return input, nil

}
