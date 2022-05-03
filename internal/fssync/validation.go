package fssync

import (
	"fmt"
	"os"
	"strings"
)

func FolderExist(path string) bool {

	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return false
}

func CreatFolderIfNotExist(path string) {
	if !FolderExist(path) {
		os.MkdirAll(path, 0755)
	}
}

func Scan(path string, base string) []string {
	list := make([]string, 0)

	dirEntries, _ := os.ReadDir(path)
	for _, entry := range dirEntries {
		if entry.IsDir() {
			list = append(list, Scan(path+"/"+entry.Name(), base)...)
		} else {
			relative := strings.Replace(path, base, "", 1)
			list = append(list, relative+"/"+entry.Name())
		}
	}

	return list
}

func Copy(source string, target string) {
	fmt.Println(source, target)
}
