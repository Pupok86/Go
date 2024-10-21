package main

import (
	"fmt"
	"io"
	"io/fs"
	"strings"

	"os"
	"path/filepath"
	// "strings"
)

func dirTree(output io.Writer, dir_path string, printFile bool) error {
	err := filepath.Walk(dir_path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		//fmt.Fprintf(output, "%d\n", tab)
		if info.Name() == filepath.Base(dir_path) {
			return nil
		}
		tab := strings.Count(path, "\\")
		if printFile {
			if info.IsDir() {
				for i := 1; i <= tab; i++ {
					if i == tab {
						fmt.Fprintf(output, "├───")
					} else {
						fmt.Fprintf(output, "│   ")
					}
				}
				fmt.Fprintf(output, "%s\n", info.Name())
			}
		} else {
			//fmt.Fprintf(output, "├───%q\n", filepath.Base(path))
			for i := 1; i <= tab; i++ {
				if i == tab {
					fmt.Fprintf(output, "├───")
				} else {
					fmt.Fprintf(output, "│   ")
				}
			}
			fmt.Fprintf(output, "%s\n", info.Name())
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", "asd", err)
	}
	return err
}
func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}

}
