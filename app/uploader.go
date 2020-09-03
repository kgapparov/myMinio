package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/minio/minio-go"
	"github.com/unsmoker/myminio/config"
)

var (
	configPath  string
	filePath    string
	fileName    string
	bucketName  string
	contentType string
)

func init() {
	flag.StringVar(&configPath, "configPath", "config/config.toml", "path to config file")
	flag.StringVar(&filePath, "filePath", ".", "path to file upload")
	flag.StringVar(&fileName, "fileName", "test", "file name to upload")
	flag.StringVar(&bucketName, "bucketName", "mybucket", "name of bucket")
}
func main() {
	flag.Parse()
	// Creating Configuration parameters from toml file
	config := config.New()
	toml.DecodeFile(configPath, config)
	fmt.Println(config.AccessKey)

	// Initialize minio client object.
	minioClient, err := minio.New(config.EndPoint, config.AccessKey, config.SecretKey, config.UseSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	objectName := fileName
	filePath := filePath + "/" + fileName

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(contentType)

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	fmt.Println(fileName)
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
