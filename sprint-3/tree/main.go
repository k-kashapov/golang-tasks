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
	"strings"
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

func doReadTree(out io.Writer, path string, printFiles bool, prefix *strings.Builder) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	os.Chdir(path)
	prefix.WriteString(BRANCHING_TRUNK)

	for _, file := range files {
		stat, err := os.Lstat(file.Name())
		if err != nil {
			return err
		}

		fmt.Fprint(out, prefix.String())

		mode := stat.Mode()

		if mode.IsRegular() {
			if printFiles {
				fmt.Fprintln(out, file.Name())
			}
		} else {
			fmt.Fprintln(out, file.Name())
			doReadTree(out, file.Name(), printFiles, prefix)
		}
	}

	os.Chdir("..")
	prefix.Reset()

	return err
}

// dirTree: `tree` program implementation, top-level function, signature is fixed.
// Write `path` dir listing to `out`. If `prinFiles` is set, files is listed along with directories.
func dirTree(out io.Writer, path string, printFiles bool) error {
	// Function to implement, signature is given, don't touch it.
	var prefix strings.Builder
	err := doReadTree(out, path, printFiles, &prefix)
	return err
}
