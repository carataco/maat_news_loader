package extract

import (
	"testing"

	"github.com/carataco/maat_news_loader/internal/types"
)

func TestNewExtractor(t *testing.T) {
	tests := []struct {
		name    string
		cfg     types.Config
		wantErr bool
	}{
		{"S£ source", types.Config{SourceType: "s3"}, false},
		{"Unknown source", types.Config{SourceType: "missing_source"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extractor, err := NewExtractor(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewExtractor() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && extractor == nil {
				t.Errorf("NewExtractor() returned nil extractor for %v", tt.cfg.SourceType)
			}
		})
	}
}
