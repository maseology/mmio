package mmio

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// CollectFilesExt returns a list of files of a given extension from a directory.
// directories should end with "/" and extensions start with ".".
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
		log.Fatal(err)
		return 0, false
	}
}
