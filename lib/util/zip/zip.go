package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

//goland:noinspection GoNameStartsWithPackageName
type ZipFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// CreateZipArchiveFileWithIgnored create a zip archive file from files, and return all ignored files
func CreateZipArchiveFileWithIgnored(savaPath string, files []ZipFile, ignoreError bool) (ignoredFiles []ZipFile, err error) {
	defer func() {
		if err != nil {
			_ = os.Remove(savaPath)
		}
	}()

	create, err := os.Create(savaPath)
	if err != nil {
		return ignoredFiles, fmt.Errorf("can not create destination zip archive file: %s", err.Error())
	}

	defer func(saveAsFile *os.File) {
		_ = saveAsFile.Close()
	}(create)

	return CreateZipArchive(create, files, ignoreError)
}

func CreateZipArchiveFile(savaPath string, files []ZipFile, ignoreError bool) error {
	create, err := os.Create(savaPath)
	if err != nil {
		return fmt.Errorf("can't create destination zip archive file: %s\n", err.Error())
	}
	_, err = CreateZipArchive(create, files, ignoreError)
	defer func(create *os.File) {
		err := create.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(create)

	return err
}

func CreateZipArchive(write io.Writer, files []ZipFile, ignoreError bool) ([]ZipFile, error) {
	zipWriter := zip.NewWriter(write)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(zipWriter)

	ignoredFiles := make([]ZipFile, 0)
	for _, file := range files {
		if err := addFileToZipArchive(zipWriter, file); err != nil {
			ignoredFiles = append(ignoredFiles, file)
			if !ignoreError {
				return ignoredFiles, err
			}
		}
	}
	return ignoredFiles, nil
}

// addFileToZipArchive add a file to zip archive
func addFileToZipArchive(zipWriter *zip.Writer, file ZipFile) error {

	zipFile, err := os.Open(file.Path)
	if err != nil {
		return fmt.Errorf("can not open %s: %s", file.Path, err.Error())
	}

	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(zipFile)

	info, err := zipFile.Stat()
	if err != nil {
		return fmt.Errorf("can not get file state for %s: %s", file.Path, err.Error())
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("create zip file header failed for %s: %s", file.Path, err.Error())
	}

	header.Name = file.Name
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("create header failed for %s: %s", file.Path, err.Error())
	}

	_, err = io.Copy(writer, zipFile)
	if err != nil {
		return fmt.Errorf("copy file content failed for %s: %s", file.Path, err.Error())
	}

	return nil
}
