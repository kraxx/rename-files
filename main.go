package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArgs(args []string) (path, renameValue string) {
	switch len(args) {
	case 1: // No args; path set to "."; rename value set to executable's
		split := strings.Split(args[0], "/")            // Remove parent dir names
		split = strings.Split(split[len(split)-1], ".") // Remove .filedescriptor
		return ".", split[0]
	case 3: // set path and rename value according to args
		return strings.TrimRight(args[1], "/"), args[2]
	default:
		fmt.Fprintln(os.Stderr, "Error: Please provide a [path string] and [rename string]")
		os.Exit(69)
	}
	return
}

func checkDir(path string) *os.File {
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

func constructNewName(num int, oldName, renameValue string) string {
	split := strings.SplitN(oldName, ".", 2)
	if len(split) == 1 {
		return renameValue + "_" + strconv.Itoa(num)
	}
	return renameValue + "_" + strconv.Itoa(num) + "." + split[1]
}

func renameFiles(dir *os.File, path, renameValue string) {
	files, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		oldName := file.Name()
		newName := constructNewName(i+1, oldName, renameValue)
		if err = os.Rename(path+"/"+oldName, path+"/"+newName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func main() {
	path, renameValue := parseArgs(os.Args)
	dir := checkDir(path)
	defer dir.Close()
	renameFiles(dir, path, renameValue)
}
