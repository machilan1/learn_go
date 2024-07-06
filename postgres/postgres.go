package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	*SnippetStore
}

// A Store factory func
func NewStore(databaseSrc string) (*Store, error) {
	db, err := sqlx.Open("postgres", databaseSrc)
	if err != nil {
		return nil, err
	}
	return &Store{
		SnippetStore: &SnippetStore{DB: db},
	}, nil

}
