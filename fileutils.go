package fileutils
import (
	"os"
	"path/filepath"
	"io"
	"io/ioutil"
	"fmt"
	"strings"
)

// CopyFile copies src file to dst. Destination file will get
// file mode specified by perm parameter
func CopyFile(dst, src string, perm os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	tmp, err := ioutil.TempFile(filepath.Dir(dst), "")
	if err != nil {
		return err
	}
	_, err = io.Copy(tmp, in)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return err
	}
	if err = tmp.Close(); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	if err = os.Chmod(tmp.Name(), perm); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	return os.Rename(tmp.Name(), dst)
}

// CopyFile recursively copies src directory to dst locatiion.
// All destination files will get file mode specified by perm parameter
func CopyDir(dst, src string, perm os.FileMode) error {
	// Check if src exists
	return filepath.Walk(src, func(path string, info os.FileInfo, _ error) error {
		// what is the new path?
		var fileType string
		if info.IsDir() {
			fileType = "DIR"
		} else {
			fileType = "FILE"
		}
		fmt.Printf("%4v: %v\n", fileType, strings.TrimPrefix(path, src))
		return nil
	})
}

