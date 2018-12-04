package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir := "."
	if len(os.Args) == 2 {
		dir = os.Args[1]
	}

	worker(dir)
}

func worker(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	fmt.Println(dir)

	for _, file := range files {
		if file.IsDir() {
			worker(filepath.Join(dir, file.Name()))
		}
	}
}
