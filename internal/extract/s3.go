package extract

import "github.com/carataco/maat_news_loader/internal/types"

// "bytes"
// "context"

// "github.com/aws/aws-sdk-go-v2/aws"
// "github.com/aws/aws-sdk-go-v2/config"
// "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
// "github.com/aws/aws-sdk-go-v2/service/s3"

type S3Extractor struct {
	s3Bucket string
	// S3Extract S3Extract
	// S3FileExtractor ExtractFunc
	// AWSS3Config func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error)
}

func NewS3Extractor(bucket string) *S3Extractor {
	extractor := &S3Extractor{s3Bucket: bucket}
	// extractor.S3Extract = extractor
	// extractor.S3FileExtractor = ExtractS3File
	// extractor.AWSS3Config = config.LoadDefaultConfig

	return extractor
}

// type S3Extract interface {
// 	Extractor(s3Bucket string, articleKey string, serializedArticle *bytes.Reader) (string, error)
// }

// type ExtractFunc func(uploadermanager *manager.Uploader, bucket string, articlekey string, article *bytes.Reader) (*manager.UploadOutput, error)

// func ExtractS3File(uploadermanager *manager.Uploader, bucket string, articlekey string, article *bytes.Reader) (*manager.UploadOutput, error) {
// 	result, err := uploadermanager.Upload(context.Background(), &s3.PutObjectInput{
// 		Bucket: aws.String(bucket),
// 		Key:    aws.String(articlekey),
// 		Body:   article,
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func (r *S3Extractor) Extractor(s3Bucket string, articleKey string, serializedArticle *bytes.Reader) (string, error) {
// 	cfg, err := r.AWSS3Config(context.TODO())
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		return "nil", err
// 	}

// 	client := s3.NewFromConfig(cfg)
// 	uploader := manager.NewUploader(client)

// 	result, err := r.S3FileExtractor(uploader, s3Bucket, articleKey, serializedArticle)

// 	if err != nil {
// 		return "nil", err
// 	}

// 	if result.Key != nil {
// 		log.Println("Uploaded object key:", *result.Key)
// 		return *result.Key, nil
// 	} else {
// 		log.Println("Object key not uploaded:", *result.Key)
// 		fmt.Println("Key not returned by S3")
// 		return *result.Key, nil
// 	}
// }

func (r *S3Extractor) ExtractS3Files() {
}

func (r *S3Extractor) Extract() ([]types.Record, error) {

	testo := []types.Record{}
	return testo, nil
}
