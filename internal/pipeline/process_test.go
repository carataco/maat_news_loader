package pipeline

import (
	"errors"
	"testing"

	"github.com/carataco/maat_news_loader/internal/extract"
	// "github.com/carataco/maat_news_loader/internal/load"
	"github.com/carataco/maat_news_loader/internal/types"
)

// ---------- Extract fakes ----------

type FakeExtractSuccess struct{}

func (r *FakeExtractSuccess) Extract(sources []string, event types.Event) ([][]any, error) {
	return [][]any{
		{'1', '5'},
		{'3', '4'},
	}, nil
}

type FakeExtractMethodFail struct{}

func (r *FakeExtractMethodFail) Extract(sources []string, event types.Event) ([][]any, error) {
	return nil, errors.New("extract method failed")
}

func FakeExtractorCtorSuccess(cfg types.Config) (extract.Extractor, error) {
	return &FakeExtractSuccess{}, nil
}

func FakeExtractorCtorFail(cfg types.Config) (extract.Extractor, error) {
	return nil, errors.New("extractor ctor failed")
}

func FakeExtractorExtractFail(cfg types.Config) (extract.Extractor, error) {
	return &FakeExtractMethodFail{}, nil
}

/// ---------- Load fakes ----------

// type FakeLoadSuccess struct{}

// func (r *FakeLoadSuccess) Load(records []types.Record) ([]string, error) {
// 	return []string{"id1", "id2"}, nil
// }

// type FakeLoadMethodFail struct{}

// func (r *FakeLoadMethodFail) Load(records []types.Record) ([]string, error) {
// 	return nil, errors.New("load method failed")
// }

// func FakeLoaderCtorSuccess(cfg types.Config) (load.Loader, error) {
// 	return &FakeLoadSuccess{}, nil
// }

// func FakeLoaderCtorFail(cfg types.Config) (load.Loader, error) {
// 	return nil, errors.New("loader ctor failed")
// }

// func FakeLoaderLoadFail(cfg types.Config) (load.Loader, error) {
// 	return &FakeLoadMethodFail{}, nil
// }

func TestRunner_Run_Failures(t *testing.T) {
	tests := []struct {
		name          string
		extractorCtor func(types.Config) (extract.Extractor, error)
		// loaderCtor    func(types.Config) (load.Loader, error)
		expectErr bool
	}{
		{
			name:          "extractor constructor fails",
			extractorCtor: FakeExtractorCtorFail,
			// loaderCtor:    FakeLoaderCtorSuccess,
			expectErr: true,
		},
		{
			name:          "extract method fails",
			extractorCtor: FakeExtractorExtractFail,
			// loaderCtor:    FakeLoaderCtorSuccess,
			expectErr: true,
		},
		// {
		// 	name:          "loader constructor fails",
		// 	extractorCtor: FakeExtractorCtorSuccess,
		// 	loaderCtor:    FakeLoaderCtorFail,
		// 	expectErr:     true,
		// },
		// {
		// 	name:          "load method fails",
		// 	extractorCtor: FakeExtractorCtorSuccess,
		// 	loaderCtor:    FakeLoaderLoadFail,
		// 	expectErr:     true,
		// },
		{
			name:          "happy path",
			extractorCtor: FakeExtractorCtorSuccess,
			// loaderCtor:    FakeLoaderCtorSuccess,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := NewRunner(types.Config{}, types.Event{})
			runner.MasterExtractor = tt.extractorCtor
			// runner.MasterLoader = tt.loaderCtor

			_, err := runner.Run()

			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
