package extract

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/carataco/maat_news_loader/internal/types"
)

type S3Extractor struct {
	s3Bucket string

	AWSS3Config func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error)
}

func NewS3Extractor(bucket string) *S3Extractor {
	extractor := &S3Extractor{s3Bucket: bucket}

	extractor.AWSS3Config = config.LoadDefaultConfig

	return extractor
}

func CreatePrefixes(startdate string, enddate string, sources []string) []string {
	prefixes := []string{}

	startdateT, _ := time.Parse("2006/01/02", startdate)
	enddateT, _ := time.Parse("2006/01/02", enddate)
	daysdiff := int(enddateT.Sub(startdateT).Hours() / 24)

	for _, source := range sources {

		if startdate == "" || enddate == "" {
			yesterday := time.Now().AddDate(0, 0, -1).Format("2006/01/02")
			prefix := "raw/" + source + "/" + yesterday + "/"
			prefixes = append(prefixes, prefix)

		} else {

			for i := range daysdiff {
				selecteddate := startdateT.Add(time.Hour * 24 * time.Duration(i)).Format("2006/01/02")
				prefix := "raw/" + source + "/" + selecteddate + "/"
				prefixes = append(prefixes, prefix)

			}
		}
	}

	return prefixes
}

func (r *S3Extractor) S3ListObjects(client *s3.Client, event types.Event, sources []string) ([]string, error) {
	objectkeys := []string{}
	prefixes := CreatePrefixes(event.StartDate, event.EndDate, sources)

	for _, prefix := range prefixes {

		paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
			Bucket: &r.s3Bucket,
			Prefix: &prefix,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(context.TODO())
			if err != nil {
				return nil, err
			}
			for _, obj := range page.Contents {

				objectkeys = append(objectkeys, *obj.Key)
			}
		}

	}

	return objectkeys, nil
}

func (r *S3Extractor) S3GetObjects(client *s3.Client, objectkeys []string) ([][]byte, error) {
	objects := [][]byte{}

	for _, obj := range objectkeys {

		result, _ := client.GetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(r.s3Bucket),
			Key:    aws.String(obj),
		})

		bodyBytes, _ := io.ReadAll(result.Body)

		objects = append(objects, bodyBytes)

		defer result.Body.Close()
	}
	return objects, nil
}

func (r *S3Extractor) Extract(sources []string, event types.Event) ([][]byte, error) {
	cfg, err := r.AWSS3Config(context.TODO())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	objectkeys, err := r.S3ListObjects(client, event, sources)
	if err != nil {
		return nil, err
	}

	objects, err := r.S3GetObjects(client, objectkeys)

	return objects, nil

}
