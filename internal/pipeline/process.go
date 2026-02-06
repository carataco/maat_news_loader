package pipeline

import (
	"fmt"

	"github.com/carataco/maat_news_loader/internal/extract"
	"github.com/carataco/maat_news_loader/internal/types"
)

type Runner struct {
	Config          types.Config
	MasterExtractor func(cfg types.Config) (extract.Extractor, error)
	// MasterLoader    func(cfg types.Config) (load.Loader, error)
}

func NewRunner(cfg types.Config) *Runner {
	runner := &Runner{Config: cfg}
	runner.MasterExtractor = extract.NewExtractor
	// runner.MasterLoader = load.NewLoader
	return runner
}

func (r *Runner) Run() ([]string, error) {
	extractor, err := r.MasterExtractor(r.Config)
	if err != nil {
		return nil, err
	}
	extracteditems, err := extractor.Extract()
	if err != nil {
		return nil, err
	}

	fmt.Println(extracteditems)
	mockresult := []string{}

	return mockresult, err
}
