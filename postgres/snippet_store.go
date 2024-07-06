package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SnippetStore struct {
	*sqlx.DB
}

// !這個Type跟書上不一樣，這只是暫時的，之後把他改成跟書上一樣。
// The type is here to act as type of an entity.
type Snippet struct {
	ID      int
	Content string
}

// A SnippetStore factory func.
func NewSnippetStore(db *sqlx.DB) *SnippetStore {
	return &SnippetStore{
		DB: db,
	}
}

type CreateSnippetParams struct {
	Content string
}

// Insert into snippet table
func (s *SnippetStore) Create(param CreateSnippetParams) error {

	// !這個值是暫時寫上去的，之後要記得拿掉。
	const _SNIPPET_ID = 1

	_, err := s.Exec(`insert into snippet (id,content) values ($1,$2)`, _SNIPPET_ID, param.Content)
	return err
}

// Find snippet by snippetId.
func (s *SnippetStore) FindSnippet(id int) (Snippet, error) {
	var sn Snippet
	if err := s.Get(&sn, `select id, content from snippet where id = $1`, id); err != nil {
		return Snippet{}, fmt.Errorf("u messed up")
	}
	return sn, nil
}
