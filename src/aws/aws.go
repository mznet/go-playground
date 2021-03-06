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

type Aws struct {
	AccessKey string
	Secret string
	Token string
	Region string
	Bucket string
}

func (a *Aws) SetConfig(accessKey string, secret string, token string, region string, bucket string) {
	a.AccessKey = accessKey
	a.Secret = secret
	a.Token = token
	a.Region = region
	a.Bucket = bucket
}

func (a Aws) UploadToS3(fileName string) {
	creds := credentials.NewStaticCredentials(a.AccessKey, a.Secret, a.Token)

	_, err := creds.Get()

	if err != nil {
		fmt.Println("Bad AWS Credential :%s", err)
	}

	cfg := aws.NewConfig().WithRegion(a.Region).WithCredentials(creds)

	svc := s3.New(session.New(), cfg)

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
	path := fileName

	params := &s3.PutObjectInput{
		Bucket: aws.String(a.Bucket),
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