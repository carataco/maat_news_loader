package main

import (
	"context"
	"encoding/json"
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
		Load: types.LoadConfig{
			PG: types.PostgresConfig{
				Host:   os.Getenv("PGHOST"),
				User:   os.Getenv("PGUSER"),
				Pwd:    os.Getenv("PGPWD"),
				Schema: os.Getenv("PGSCHEMA"),
				Table:  os.Getenv("PGTABLE"),
			},
		},
	}
	log.Println("Lambda initialized")

}

func LambdaHandler(ctx context.Context, event types.Event) (int64, error) {
	fmt.Println("Running Lambda Handler")
	runner := pipeline.NewRunner(cfg, event)
	results, err := runner.Run()

	if err != nil {
		return 0, err
	}

	return results, nil
}

func main() {
	if os.Getenv("LOCAL") == "true" {
		fmt.Println("Running locally")
		data, err := os.ReadFile("event_file.json")

		var event types.Event
		if err := json.Unmarshal(data, &event); err != nil {
			panic(err)
		}

		results, err := LambdaHandler(context.TODO(), event)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Rows Loaded: %+v\n", results)
		return
	}

	lambda.Start(LambdaHandler)
}
