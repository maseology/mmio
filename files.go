package mmio

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// DeleteFile deletes the specified file
func DeleteFile(fp string) {
	if _, ok := FileExists(fp); ok {
		if err := os.Remove(fp); err != nil {
			log.Fatal(err)
		}
	}
}

// DeleteAllInDirectory deletes all files of a given extension in a specified directory
// exension format: ".***"
func DeleteAllInDirectory(dir, ext string) {
for _, fp := range CollectFilesExt(dir,ext) {
	DeleteFile(fp)
}
}

// CollectFilesExt returns a list of files of a given extension from a directory.
// directories should end with "/" and extensions start with ".".
// exension format: ".***"
func CollectFilesExt(dir, ext string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var flst []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ext {
			flst = append(flst, dir+file.Name())
		}
	}
	return flst
}

// FileExists checks if a file exists and returns its size
func FileExists(fp string) (int64, bool) {
	if fi, err := os.Stat(fp); err == nil {
		return fi.Size(), true
	} else if os.IsNotExist(err) {
		return 0, false
	} else {
		// log.Fatalf("mmio.FileExists: %v", err)
		return 0, false
	}
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileName returns the file name
func FileName(fp string, withExtension bool) string {
	fn := filepath.Base(fp)
	if !withExtension {
		return strings.TrimSuffix(fn, filepath.Ext(fp))
	}
	return fn
}

// GetExtension returns the file path extension
func GetExtension(fp string) string {
	return  filepath.Ext(fp)
}

// RemoveExtension returns the file path without its extension
func RemoveExtension(fp string) string {
	return strings.TrimSuffix(fp, filepath.Ext(fp))
}
