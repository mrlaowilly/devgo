package fssync

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func FolderExist(path string) bool {

	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
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

func Copy(source string, target string) error {
	fmt.Println(source, target)

	fi, err := os.Open(source)
	defer fi.Close()
	if err != nil {
		return err
	}

	parentFolder := path.Dir(target)
	err = os.MkdirAll(parentFolder, 0755)
	if err != nil {
		return err
	}

	fo, err := os.Create(target)
	defer fo.Close()
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	for {
		n, err := fi.Read(buffer)

		if n == 0 {
			break
		} else if err != nil {
			return err
		}

		if _, err := fo.Write(buffer[:n]); err != nil {
			break
		}
	}

	return nil
}
