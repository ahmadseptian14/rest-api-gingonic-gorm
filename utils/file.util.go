package utils

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	rand.Seed(time.Now().UnixMilli())
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	log.Println("Content-Type", contentType)
	result := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			result = true
			break
		}
	}

	return result
}

func FileValidationByExtention(fileHeader *multipart.FileHeader, fileExtention []string) bool {
	extention := filepath.Ext(fileHeader.Filename)
	log.Println("extension", extention)
	result := false

	for _, typeFile := range fileExtention {
		if extention == typeFile {
			result = true
			break
		}
	}

	return result
}

func RandomFileName(extentionFile string, prefix ...string) string {
	currentPrefix := "file"
	if len(prefix) > 0 {
		if prefix[0] != "" {
			currentPrefix  = prefix[0]
		}
	}

	currentTime := time.Now().UTC().Format("20061206")
	filename := fmt.Sprintf("%s-%s%s%s", currentPrefix, currentTime, RandomString(10), extentionFile)

	return filename
}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filename string) bool {
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))
	if errUpload != nil {
		log.Println("Cant save file")
		return false
	}else{
		return true
	}
}

func RemoveFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("file not found")
		}
		log.Println("Failed to remove file")
		return err
	}
	return nil
}
