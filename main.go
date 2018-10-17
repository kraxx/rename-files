package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

var DIR_SEPARATOR string = "/"

func parseArgs(args []string) (path, renameValue string) {
	switch len(args) {
	case 1: // No args; path set to "."; rename value set to executable's
		split := strings.Split(args[0], DIR_SEPARATOR)  // Remove parent dir names
		split = strings.Split(split[len(split)-1], ".") // Remove .extension
		return ".", split[0]
	case 3: // set path and rename value according to args
		return strings.TrimRight(args[1], DIR_SEPARATOR), args[2]
	default:
		fmt.Fprintln(os.Stderr, "Error: Please provide a [path string] and [rename string]; OR no args")
		os.Exit(1)
	}
	return
}

func checkPathIsDir(path string) *os.File {
	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	if fs, err := dir.Stat(); err != nil {
		panic(err)
	} else if !fs.IsDir() {
		panic(errors.New("Error: Path is not a directory"))
	}
	return dir
}

func removeExecutableFromList(files []os.FileInfo) []os.FileInfo {
	exe, err := os.Open(os.Args[0])
	defer exe.Close()
	if err != nil {
		panic(err)
	}
	exeStat, err := exe.Stat()
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		if os.SameFile(file, exeStat) {
			return append(files[:i], files[i+1:]...)
		}
	}
	return files
}

func processFileNameForNumericSort(a, b string) (string, string) {
	aSplit, bSplit := strings.Split(a, "."), strings.Split(b, ".")
	if aSplit[0] == bSplit[0] {
		return aSplit[1], bSplit[1]
	}
	return aSplit[0], bSplit[0]
}

func sortFilesNumerically(files []os.FileInfo) {
	sort.Slice(files, func(x, y int) bool {
		a, b := files[x].Name(), files[y].Name()
		a, b = processFileNameForNumericSort(a, b)
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				aInt, aErr := strconv.Atoi(a[i:])
				bInt, bErr := strconv.Atoi(b[i:])
				if aErr != nil || bErr != nil {
					return a < b
				}
				return aInt < bInt
			}
		}
		return true
	})
}

func getFilesInDirectory(dir *os.File) []os.FileInfo {
	files, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	files = removeExecutableFromList(files)
	sortFilesNumerically(files)
	return files
}

func constructNewName(num int, oldName, renameValue string) string {
	split := strings.SplitN(oldName, ".", 2)
	if len(split) == 1 {
		return renameValue + "_" + strconv.Itoa(num)
	}
	return renameValue + "_" + strconv.Itoa(num) + "." + split[1]
}

func renameFiles(files []os.FileInfo, path, renameValue string) {
	for i, file := range files {
		oldName := file.Name()
		newName := constructNewName(i+1, oldName, renameValue)
		if err := os.Rename(path+DIR_SEPARATOR+oldName, path+DIR_SEPARATOR+newName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func init() {
	if runtime.GOOS == "windows" {
		DIR_SEPARATOR = "\\"
	}
}

func main() {
	path, renameValue := parseArgs(os.Args)
	dir := checkPathIsDir(path)
	defer dir.Close()
	files := getFilesInDirectory(dir)
	renameFiles(files, path, renameValue)
}
