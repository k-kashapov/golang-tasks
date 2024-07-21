package main

/*
Course `Web services on Go`, week 1, homework, `tree` program.
See: week_01\materials.zip\week_1\99_hw\tree

mkdir -p week01_homework/tree
pushd week01_homework/tree
go mod init tree
go mod tidy
pushd ..
go work init
go work use ./tree/
go vet tree
gofmt -w tree
go test -v tree
go run tree . -f
go run tree ./tree/testdata
cd tree && docker build -t mailgo_hw1 .

https://en.wikipedia.org/wiki/Tree_(command)
https://mama.indstate.edu/users/ice/tree/
https://stackoverflow.com/questions/32151776/visualize-tree-in-bash-like-the-output-of-unix-tree

*/

import (
	"fmt"
	"io"
	"os"
)

/*
	Example output:

	├───project
	│	└───gopher.png (70372b)
	├───static
	│	├───a_lorem
	│	│	├───dolor.txt (empty)
	│	├───css
	│	│	└───body.css (28b)
	...
	│			└───gopher.png (70372b)

	- path should point to a directory,
	- output all dir items in sorted order, w/o distinction file/dir
	- last element prefix is `└───`
	- other elements prefix is `├───`
	- nested elements aligned with one tab `	` for each level
*/

const (
	EOL             = "\n"
	BRANCHING_TRUNK = "├───"
	LAST_BRANCH     = "└───"
	TRUNC_TAB       = "│\t"
	LAST_TAB        = "\t"
	EMPTY_FILE      = "empty"
	ROOT_PREFIX     = ""

	USE_RECURSION_ENV_KEY = "RECURSIVE_TREE"
	USE_RECURSION_ENV_VAL = "YES"
)

func main() {
	// This code is given
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage: go run main.go . [-f]")
	}

	out := os.Stdout
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func findLastDir(files []os.DirEntry) int {
	last := -1
	for i, file := range files {
		if file.IsDir() {
			last = i
		}
	}

	return last
}

func doReadDir(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	os.Chdir(path)

	var isLast bool
	lastDir := findLastDir(files)
	total := len(files)

	for i, file := range files {
		if !printFiles && !file.IsDir() {
			continue
		}

		if printFiles {
			isLast = i+1 == total
		} else {
			isLast = i == lastDir
		}

		fmt.Fprint(out, prefix)

		var next_prefix string

		if isLast {
			fmt.Fprint(out, LAST_BRANCH)
			next_prefix = prefix + LAST_TAB
		} else {
			fmt.Fprint(out, BRANCHING_TRUNK)
			next_prefix = prefix + TRUNC_TAB
		}

		fmt.Fprint(out, file.Name())

		if printFiles && !file.IsDir() {
			stat, _ := os.Lstat(file.Name())
			size := stat.Size()
			if size == 0 {
				fmt.Fprintln(out, " (empty)")
			} else {
				fmt.Fprintf(out, " (%db)\n", size)
			}
		} else {
			fmt.Fprintln(out, "")
		}

		doReadDir(out, file.Name(), printFiles, next_prefix)
	}

	os.Chdir("..")

	return err
}

// dirTree: `tree` program implementation, top-level function, signature is fixed.
// Write `path` dir listing to `out`. If `prinFiles` is set, files is listed along with directories.
func dirTree(out io.Writer, path string, printFiles bool) error {
	// Function to implement, signature is given, don't touch it.
	err := doReadDir(out, path, printFiles, ROOT_PREFIX)
	return err
}
