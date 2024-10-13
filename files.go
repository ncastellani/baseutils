package baseutils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ListFolderFiles
// take a folder path and walk through each item on it
func ListFolderFiles(folder string, includeDirectories bool) (files []os.FileInfo, err error) {

	// walk through each file on the folder
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// filter out the folders
		if !includeDirectories && info.IsDir() {
			return nil
		}

		// append this file/folder into the list
		files = append(files, info)

		return nil
	})

	return
}

// CopyFile
// open and read the source file contents, then, call WriteFileto write the new file
func CopyFile(source, dest string) (err error) {

	// read the file contents
	file, err := os.Open(source)
	if err != nil {
		return
	}

	defer file.Close()

	// write the file into the new path
	contents, err := io.ReadAll(file)
	if err != nil {
		return
	}

	return WriteFile(dest, &contents)
}

// WriteFile
// take an path to write into and content in bytes and write the file creating the folder tree recursively
func WriteFile(path string, content *[]byte) (err error) {

	// determine the folder structure
	pathSplitted := strings.Split(filepath.ToSlash(path), "/")
	folder := strings.Join(pathSplitted[:len(pathSplitted)-1], "/")

	// create the folder if not exists
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// open the destination file
	destFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}

	defer destFile.Close()

	// write into the file
	_, err = destFile.Write(*content)
	if err != nil {
		return
	}

	return
}
