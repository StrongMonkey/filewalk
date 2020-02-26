package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("need exactly two arguments")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	for _, r := range Walk(file) {
		fmt.Println(r)
	}
}

func Walk(f *os.File) []string {
	var r []string
	base := basename(f)
	info, err := os.Stat(f.Name())
	if err != nil {
		panic(err)
	}
	if info.IsDir() {
		files, err := f.Readdir(-1)
		if err != nil {
			panic(err)
		}
		for _, cf := range files {
			childFile, err := os.Open(filepath.Join(f.Name(), cf.Name()))
			if err != nil {
				panic(err)
			}
			r1 := Walk(childFile)
			for i, s := range r1 {
				r1[i] = base + "/" + s
			}
			r = append(r, r1...)
		}
	} else {
		r = append(r, basename(f))
	}
	return r
}

func basename(f *os.File) string {
	return filepath.Base(f.Name())
}
