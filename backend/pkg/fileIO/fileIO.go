package fileIO

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func CreateTempFolder(name string) string {
	dir, err := os.MkdirTemp("temp", name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return dir
}

func CreateJsonFormatBytes(path string, data interface{}) []byte {
	file, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

func CreateTempFile(path string, file *multipart.FileHeader) (*os.File, string) {
	var contentType = strings.Split(file.Header.Get("Content-Type"), "/")[1]
	fileName := fmt.Sprintf("%s.%s", file.Filename, contentType)
	dst, err := os.Create(fmt.Sprintf("%s/%s", path, fileName))
	if err != nil {
		log.Fatal(err.Error())
	}
	return dst, fileName
}

func CopyFile(dst io.Writer, src io.Reader) error {
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func AddToFolder(path string, importFile *multipart.FileHeader) (string, error) {
	src, err := importFile.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	dst, fileName := CreateTempFile(path, importFile)
	defer dst.Close()
	err = CopyFile(dst, src)
	return fileName, nil
}

func RemoveAll(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err.Error())
	}
}
