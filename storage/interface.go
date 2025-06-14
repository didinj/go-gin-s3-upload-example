package storage

import (
	"io"
)

var UploadFile func(string, io.Reader) (string, error)
