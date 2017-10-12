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
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AccessKey string `yaml:"access_key"`
	Secret string `yaml:"secret"`
	Token string `yaml: "token"`
	Region string `yaml: "region"`
	Bucket string `yaml: "token"`
}

func UploadToS3(fileName string) {
	var config Config
	filepath := ".aws/credential.yaml"

	source, readErr := ioutil.ReadFile(filepath)

	if readErr != nil {
		panic(readErr)
	}

	readErr = yaml.Unmarshal(source, &config)
	if readErr != nil {
		panic(readErr)
	}
	fmt.Printf("Value: %#v\n", config)

	creds := credentials.NewStaticCredentials(config.AccessKey, config.Secret, config.Token)

	_, err := creds.Get()

	if err != nil {
		fmt.Println("Bad AWS Credential :%s", err)
	}

	cfg := aws.NewConfig().WithRegion(config.Region).WithCredentials(creds)

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
		Bucket: aws.String(config.Bucket),
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