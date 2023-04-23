package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <directory>\n", os.Args[0])
		os.Exit(1)
	}

	root := os.Args[1]
	tree(root, "")
}

func tree(path string, prefix string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	for i, file := range files {
		name := file.Name()
		if name == ".git" {
			continue
		}
		if name == "node_modules" {
			continue
		}
		if name == ".vscode" {
			continue
		}

		if file.IsDir() {
			fmt.Printf("%s├───%s/\n", prefix, name)
			if i == len(files)-1 {
				tree(filepath.Join(path, name), prefix+"    ")
			} else {
				tree(filepath.Join(path, name), prefix+"│   ")
			}
		} else {
			if i == len(files)-1 {
				fmt.Printf("%s└───%s\n", prefix, name)
			} else {
				fmt.Printf("%s├───%s\n", prefix, name)
			}
		}
	}
}
