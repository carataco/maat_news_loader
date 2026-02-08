package load

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type PGLoader struct {
	Host   string
	User   string
	Pwd    string
	Schema string
	Table  string
}

func NewPGLoader(host string, user string, pwd string, schema string, table string) *PGLoader {
	loader := &PGLoader{Host: host, User: user, Pwd: pwd, Schema: schema, Table: table}

	return loader
}

func (r *PGLoader) Load(extractedobjects [][]any) (int64, error) {
	dsn := fmt.Sprintf(r.Host, r.User, r.Pwd)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return 0, err
	}

	fmt.Println("Number of rows to insert:", len(extractedobjects))
	copyCount, err := conn.CopyFrom(
		context.Background(),
		pgx.Identifier{r.Schema, r.Table},
		[]string{"payload"},
		pgx.CopyFromRows(extractedobjects),
	)

	if err != nil {
		return 0, err
	}

	defer conn.Close(context.Background())

	return copyCount, nil
}
