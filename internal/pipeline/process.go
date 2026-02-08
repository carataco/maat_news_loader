package pipeline

import (
	"github.com/carataco/maat_news_loader/internal/extract"
	"github.com/carataco/maat_news_loader/internal/load"
	"github.com/carataco/maat_news_loader/internal/types"
)

type Runner struct {
	Config          types.Config
	Event           types.Event
	MasterExtractor func(cfg types.Config) (extract.Extractor, error)
	MasterLoader    func(cfg types.Config) (load.Loader, error)
}

func NewRunner(cfg types.Config, event types.Event) *Runner {

	runner := &Runner{Config: cfg, Event: event}
	runner.MasterExtractor = extract.NewExtractor
	runner.MasterLoader = load.NewLoader
	return runner
}

func (r *Runner) Run() (int64, error) {

	extractor, err := r.MasterExtractor(r.Config)
	if err != nil {
		return 0, err
	}

	extractedobjects, err := extractor.Extract(r.Config.SourceIDS, r.Event)
	if err != nil {
		return 0, err
	}

	loader, err := r.MasterLoader(r.Config)
	if err != nil {
		return 0, err
	}

	rowsloaded, err := loader.Load(extractedobjects)

	return rowsloaded, err
}
