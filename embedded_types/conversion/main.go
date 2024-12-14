package main

import "fmt"

type (
	Backend interface {
		RemoteStateBlock() string
	}

	S3 interface {
		Key() string
	}
)

type AWSS3 struct {
	bucket string
}

type AzureS3 struct {
	AWSS3
}

func (a AWSS3) RemoteStateBlock() string {
	return fmt.Sprintf("terraform { bucket = %s }", a.bucket)
}

func (a AWSS3) Key() string {
	return "my-aws-key"
}

func (a AzureS3) Key() string {
	return "my-azure-key"
}

func main() {
	var backend Backend = AWSS3{bucket: "my-aws-bucket"}
	// var backend Backend = AzureS3{AWSS3: AWSS3{bucket: "my-azure-bucket"}}
	fmt.Println(backend.RemoteStateBlock())

	s3, ok := backend.(S3)
	if !ok {
		fmt.Println("Not an S3 backend")
	}
	fmt.Println(s3.Key())

	// if a, ok := interface{}(azure).(AzureS3); ok {
	// 	fmt.Println(a.Key())
	// }
}
