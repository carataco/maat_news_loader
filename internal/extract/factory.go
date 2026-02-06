package extract

import (
	"fmt"

	"github.com/carataco/maat_news_loader/internal/types"
)

func NewExtractor(cfg types.Config) (Extractor, error) {
	if cfg.SourceType == "s3" {
		fmt.Println("Extracting from s3")
		return NewS3Extractor(cfg.Extract.S3.Bucket), nil
	} else {
		return nil, fmt.Errorf("unknown source type")
	}

}
