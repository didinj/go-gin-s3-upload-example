package config

import "os"

var (
	StorageType = os.Getenv("STORAGE_TYPE") // "s3" or "local"

	S3Region    = os.Getenv("AWS_REGION")
	S3Bucket    = os.Getenv("AWS_BUCKET")
	S3AccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	S3SecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	UploadPath = "./uploads" // for local storage
)
