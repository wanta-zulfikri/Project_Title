package helper

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

func Upload(c echo.Context, credential_id, credential_access string) (string, error) {
	file, err := c.FormFile("file") 
	if err != nil {
		return "", err 
	} 
	src, err := file.Open()
	if err != nil {
		return "",err 
	}
	defer src.Close() 
	result, err := UploadToS3(c, file.Filename, src, credential_id, credential_access)
	if err != nil {
		return "",err 
	}
	
	return result, nil  
}

// ======================================================================
// UPLOAD TO S3
// ======================================================================
// Helper
func UploadToS3(c echo.Context, filename string, src multipart.File, credential_id, credential_access string) (string, error) {
	logger := c.Logger()
	s3Config := &aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(credential_id, credential_access, ""),
	}
	s3Session, _ := session.NewSession(s3Config)
	uploader := s3manager.NewUploader(s3Session) 
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("AWSBUCKET/"),
		Key: aws.String(filename), 
		Body: src,
		ACL: aws.String("public-read"),
	})  
	if err != nil {
		logger.Fatal(err)
		return "", err 
	}
	return result.Location, nil 
}
