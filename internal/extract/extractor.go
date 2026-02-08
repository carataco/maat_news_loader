package extract

import "github.com/carataco/maat_news_loader/internal/types"

type Extractor interface {
	Extract(sources []string, event types.Event) ([][]byte, error)
}
