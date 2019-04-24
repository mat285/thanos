package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		printUsage()
	}
	if strings.ToLower(os.Args[1]) != "snap" {
		printUsage()
	}
	dir := os.Args[2]
	if len(dir) == 0 {
		printUsage()
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if !info.IsDir() {
			return snap(path, info)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func snap(path string, info os.FileInfo) error {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	outlines := make([]byte, 0, len(contents))

	keep := true
	for _, b := range contents {
		if keep {
			outlines = append(outlines, b)
		}
		if b == '\n' {
			keep = !keep
		}
	}
	return ioutil.WriteFile(path, outlines, info.Mode())
}

func printUsage() {
	fmt.Println("Usage:\n\nthanos snap [dir]")
	os.Exit(0)
}
