package aws

import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"os"
	"bytes"
	"net/http"
)

func UploadToS3(fileName) {
	aws_access_key_id := "Insert Key ID here"
	aws_secret_access_key := "Insert Secret Here"
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	bucketName := "test_bucket"

	_, err := creds.Get()

	if err != nil {
		fmt.Println("Bad AWS Credential :%s", err)
	}

	cfg := aws.NewConfig().WithRegion("eu-central-1").WithCredentials(creds)

	svc := s3.New(sessions.New(), cfg)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("err opening file: %s", err)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	size := fileInfo.Size()

	buffer := make([] byte, size)

	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path = fileName

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(path),
		Body: fileBytes,
		ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	}

	resp, err := svc.PutObject(params)
	if err != nil {
		fmt.Printf("bad response: %s", err)
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))

}