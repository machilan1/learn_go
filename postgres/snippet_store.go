package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SnippetStore struct {
	*sqlx.DB
}

// TODO: 這個Type跟書上不一樣，這只是暫時的，記得把他改成正常的。
type Snippet struct {
	ID      int
	Content string
}

type CreateSnippetParams struct {
	Content string
}

func NewSnippetStore(db *sqlx.DB) *SnippetStore {
	return &SnippetStore{
		DB: db,
	}
}

// TODO: 這個值是暫時寫上去的，之後要記得拿掉。
const MAGICAL_TEMP_ID = 1

func (s *SnippetStore) Create(param CreateSnippetParams) error {
	_, err := s.Exec(`insert into snippet (id,content) values ($1,$2)`, MAGICAL_TEMP_ID, param.Content)
	return err
}

func (s *SnippetStore) FindSnippet(id int) (Snippet, error) {
	var sn Snippet
	if err := s.Get(&sn, `select id, content from snippet where id = $1`, id); err != nil {
		return Snippet{}, fmt.Errorf("u messed up")
	}
	return sn, nil
}
