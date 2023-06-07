package helper

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/labstack/echo/v4"
    
)

type UploadResult struct{
	Path string `json:"path" xml:"path"`
}

func upload(c echo.Context) error {
	file, err := c.FormFile("file") 
	if err != nil {
		return err 
	} 
	src, err := file.Open()
	if err != nil {
		return err 
	}
	defer src.Close() 
	result, err := UploadToS3(c, file.Filename, src)
	if err != nil {
		return err 
	}
	data := &UploadResult {
		Path: result,
	}
	return c.JSON(http.StatusOK,data)
}

// ======================================================================
// UPLOAD TO S3
// ======================================================================
// Helper
func UploadToS3(c echo.Context, filename string, src multipart.File) (string, error) {
	logger := c.Logger()
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Fatal(err)
		return "", err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client) 
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("AWSBUCKET/"),
		Key: aws.String(filename), 
		Body: src,
	})  
	if err != nil {
		logger.Fatal(err)
		return "", err 
	}
	return result.Location, nil 
}
