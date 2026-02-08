package load

import (
	"fmt"

	"github.com/carataco/maat_news_loader/internal/types"
)

func NewLoader(cfg types.Config) (Loader, error) {
	if cfg.TargetType == "postgres" {
		fmt.Println("Loading into Postgres")
		return NewPGLoader(cfg.Load.PG.Host, cfg.Load.PG.User, cfg.Load.PG.Pwd, cfg.Load.PG.Schema, cfg.Load.PG.Table), nil
	} else {
		return nil, fmt.Errorf("unknown target type")
	}

}
