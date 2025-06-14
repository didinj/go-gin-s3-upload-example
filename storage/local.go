package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/yourusername/go-upload-api/config"
)

func init() {
	if config.StorageType == "local" {
		UploadFile = uploadLocal
	}
}

func uploadLocal(filename string, file io.Reader) (string, error) {
	path := filepath.Join(config.UploadPath, filename)
	os.MkdirAll(config.UploadPath, os.ModePerm)

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/uploads/%s", filename), nil
}
