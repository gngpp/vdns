package file

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"vdns/lib/vlog"
)

var log = vlog.Default()

// IsDir Determine whether it is a directory
func IsDir(path string) bool {
	open, err := os.Open(path)
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Error(err)
		}
	}(open)

	if err != nil || os.IsNotExist(err) {
		return false
	}

	stat, err := open.Stat()
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// IsFile Determine whether it is a file
func IsFile(path string) bool {
	return !IsDir(path)
}

// CurrentDir get the current file dir of the caller
func CurrentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}

// SavaToFile Save the byte array to the file
func SavaToFile(filePath string, bytes []byte) error {
	return ioutil.WriteFile(filePath, bytes, fs.ModePerm)
}

// MakeDir Recursively create new folders based on the incoming folder name
func MakeDir(fileDir string) error {
	return os.MkdirAll(fileDir, fs.ModePerm)
}

// MakeDirByFile create a folder recursively based on the incoming file name
// ./dir/filename  /home/dir/filename
func MakeDirByFile(filePath string) error {
	temp := strings.Split(filePath, "/")
	if len(temp) <= 2 {
		return errors.New("please input complete file name like ./dir/filename or /home/dir/filename")
	}
	dirPath := strings.Join(temp[0:len(temp)-1], "/")
	return MakeDir(dirPath)
}

// Exist determine whether the file exists
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//goland:noinspection ALL
func Size(path string) int64 {
	stat, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return stat.Size()
}

// InsertSuffix insert a suffix to filepath
func InsertSuffix(src string, suffix string) string {
	ext := path.Ext(src)
	return fmt.Sprintf("%s%s%s", src[:len(src)-len(ext)], suffix, ext)
}

// ReplaceExt replace ext
func ReplaceExt(src string, ext string) string {
	srcExt := path.Ext(src)
	return fmt.Sprintf("%s%s", src[:len(src)-len(srcExt)], ext)
}

// Read file contents
func Read(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	return string(file), err
}

func Create(first string, more ...string) error {
	absPath := first
	for i := range more {
		if !strings.HasPrefix(more[i], "/") {
			more[i] = "/" + more[i]
		}
		absPath += more[i]
	}
	if !Exist(absPath) {
		err := MakeDirByFile(absPath)
		if err != nil {
			return err
		}
	}
	_, err := os.Create(absPath)
	return err
}

// DelFile  removes path and any children it contains.
func DelFile(absPath string) error {
	return os.RemoveAll(absPath)
}

//goland:noinspection ALL
func GetPathDirs(absPath string) ([]string, error) {
	dirs := make([]string, 0)
	if Exist(absPath) {
		dir, err := ioutil.ReadDir(absPath)
		if err != nil {
			return nil, err
		}
		for _, f := range dir {
			if f.IsDir() {
				dirs = append(dirs, f.Name())
			}
		}
	}
	return dirs, nil
}

//goland:noinspection ALL
func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

//goland:noinspection ALL
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2f B", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2f KiB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MiB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GiB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TiB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2f PiB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
