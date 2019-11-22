package mmio

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// DeleteFile deletes the specified file or directory
func DeleteFile(fp string) {
	if DirExists(fp) {
		if err := os.RemoveAll(fp); err != nil {
			log.Fatalf("files.go DeleteFile error 2: %v", err)
		}
	} else if _, ok := FileExists(fp); ok {
		if err := os.Remove(fp); err != nil {
			log.Fatalf("files.go DeleteFile error 1: %v", err)
		}
	}
}

// DeleteAllInDirectory deletes all files of a given extension in a specified directory
// exension format: ".***"
func DeleteAllInDirectory(dir, ext string) {
	for _, fp := range FileListExt(dir, ext) {
		DeleteFile(fp)
	}
}

// DeleteAllSubdirectories deletes all subdirectories within a specified directory
func DeleteAllSubdirectories(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			os.RemoveAll(filepath.Join(dir, f.Name()))
		}
	}
}

// FileExists checks if a file exists and returns its size
func FileExists(path string) (int64, bool) {
	if fi, err := os.Stat(path); err == nil {
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

// IsDir check if the entered path is a directory
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	m := fi.Mode()
	return m.IsDir()
}

// MakeDir checks if directory exists, if not, creates it
func MakeDir(path string) {
	if !DirExists(path) {
		// if err := os.Mkdir(path, os.ModeDir); err != nil {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

// CleanDir adds a "/" at end if it does not exist
func cleanDir(dir string) string {
	if dir[len(dir)-1:] != "/" {
		dir += "/"
	}
	return dir
}

// FileRename renames a file
func FileRename(oldName, newName string, overwrite bool) {
	if _, ok := FileExists(oldName); ok {
		err := os.Rename(oldName, newName)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// FileName returns the file name
func FileName(fp string, withExtension bool) string {
	fn := filepath.Base(fp)
	if !withExtension {
		return strings.TrimSuffix(fn, filepath.Ext(fp))
	}
	return fn
}

// FileList returns a recursive list of files from a given directory
func FileList(path string) ([]string, error) {
	s := []string{}
	err := filepath.Walk(path,
		func(fp string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				s = append(s, strings.Replace(fp, string(92), string(47), -1))
				// s = append(s, strings.Replace(strings.Replace(fp, string(92), string(47), -1), path, "", -1))
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return s, nil
}

// FileListExt returns a list of files of a given extension from a directory.
// directories should end with "/" and extensions start with ".".
// exension format: ".***"
func FileListExt(dir, ext string) []string {
	dir = cleanDir(dir)
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

// DirList returns a list of subdirectories
func DirList(root string) ([]string, error) {
	s := []string{}
	err := filepath.Walk(root,
		func(fp string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && root != fp {
				s = append(s, strings.Replace(fp, string(92), string(47), -1))
				// s = append(s, strings.Replace(strings.Replace(fp, string(92), string(47), -1), path, "", -1))
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetExtension returns the file path extension
func GetExtension(fp string) string {
	return filepath.Ext(fp)
}

// RemoveExtension returns the file path without its extension
func RemoveExtension(fp string) string {
	return strings.TrimSuffix(fp, filepath.Ext(fp))
}

// GetFileDir returns the directory that contains the given filepath
func GetFileDir(fp string) string {
	if fp[len(fp)-1:] == "/" {
		return filepath.Dir(fp[:len(fp)-1])
	}
	return filepath.Dir(fp)
}
