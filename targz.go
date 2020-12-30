package mmio

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// CompressTarGZ converts a path to a *.tar.gz
func CompressTarGZ(path string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("unable to tar files - %v", err.Error())
	}
	fps, err := FileList(path)
	if err != nil {
		return err
	}
	return compressTarGZ(fps, path)
}

// CompressTarGZext converts files (with a given extension) in a path to a *.tar.gz
func CompressTarGZext(path, ext string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("unable to tar files - %v", err.Error())
	}
	return compressTarGZ(FileListExt(path, ext), path)
}

func compressTarGZ(fps []string, path string) error {
	tgzf, err := os.Create(path + "-" + MMtime(time.Now()) + ".tar.gz")
	if err != nil {
		return err
	}
	defer tgzf.Close()

	gzw := gzip.NewWriter(tgzf)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	path = strings.Replace(path, string(92), string(47), -1)
	if path[:len(path)-1] != string(47) {
		path += string(47)
	}
	for _, fp := range fps { // from https://socketloop.com/tutorials/golang-archive-directory-with-tar-and-gzip
		f, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			return err
		}

		// prepare the tar header
		header := new(tar.Header)
		header.Name = strings.Replace(f.Name(), path, "", -1)
		header.Size = fi.Size()
		header.Mode = int64(fi.Mode())
		header.ModTime = fi.ModTime()

		if err = tw.WriteHeader(header); err != nil {
			return err
		}
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
	}
	return nil

	// // walk path // from https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
	// return filepath.Walk(path, func(file string, fi os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !fi.Mode().IsRegular() {
	// 		return nil
	// 	}

	// 	// create a new dir/file header
	// 	header, err := tar.FileInfoHeader(fi, fi.Name())
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// update the name to correctly reflect the desired destination when untaring
	// 	header.Name = strings.TrimPrefix(strings.Replace(file, path, "", -1), string(filepath.Separator))

	// 	// write the header
	// 	if err := tw.WriteHeader(header); err != nil {
	// 		return err
	// 	}

	// 	// open files for taring
	// 	f, err := os.Open(file)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// copy file data into tar writer
	// 	if _, err := io.Copy(tw, f); err != nil {
	// 		return err
	// 	}

	// 	// manually close here after each file operation; defering would cause each file close
	// 	// to wait until all operations have completed.
	// 	f.Close()

	// 	return nil
	// })
}

// ExtractTarGZ extracts a *.tar.gz file to a directory
// from https://gist.github.com/indraniel/1a91458984179ab4cf80
func ExtractTarGZ(fp string) (string, error) {
	f, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return "", err
	}
	defer gzf.Close()

	odir := strings.Replace(fp, ".tar.gz", string(47), -1)
	MakeDir(odir)

	tarReader := tar.NewReader(gzf)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		name := odir + header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			os.Mkdir(name, 0755)
		case tar.TypeReg:
			dir := GetFileDir(name)
			if !DirExists(dir) {
				MakeDir(dir)
			}
			f, err := os.Create(name)
			if err != nil {
				return "", err
			}
			io.Copy(f, tarReader)
			f.Close()
			// data := make([]byte, header.Size)
			// _, err := tarReader.Read(data)
			// if err != nil {
			// 	return err
			// }
			// ioutil.WriteFile(name, data, 0755)
		default:
			return "", fmt.Errorf("unknown file type: %c in file %s", header.Typeflag, name)
		}
	}
	return odir, nil
}
