package repoutils

import (
	"database/sql"
	_ "embed"
	// _ "modernc.org/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	*sql.DB
	Path string
}

var R Repo

func DB() *sql.DB {
	return R.DB // database
}

func (R *Repo) RunQuery1(qs QuerySpec) (any, error) {
	// e := gs.BuildQuery()
	return nil, nil
}
