package mmio

import (
	"io/ioutil"
	"log"
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
