package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"

	"github.com/carataco/maat_news_loader/internal/pipeline"
	"github.com/carataco/maat_news_loader/internal/types"
)

var cfg types.Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	cfg = types.Config{
		SourceType: os.Getenv("SOURCE_TYPE"),
		SourceIDS:  strings.Split(os.Getenv("SOURCE_IDS"), ","),
		TargetType: os.Getenv("TARGET_TYPE"),
		AWSRegion:  os.Getenv("AWS_REGION"),
		Extract: types.ExtractConfig{
			S3: types.S3Config{
				Bucket: os.Getenv("S3_BUCKET"),
			},
		},
	}
	log.Println("Lambda initialized")

}

func LambdaHandler() ([]string, error) {
	fmt.Println("Running Lambda Handler")

	runner := pipeline.NewRunner(cfg)
	results, err := runner.Run()

	if err != nil {
		return nil, err
	}

	return results, nil
}

func main() {
	if os.Getenv("LOCAL") == "true" {
		fmt.Println("Running locally")
		results, err := LambdaHandler()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Articles Loaded: %+v\n", len(results))
		return
	}

	lambda.Start(LambdaHandler)
}
