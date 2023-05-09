package dictionnary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

type Dictionnary struct {
	db *badger.DB
}

type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

func New(dir string) (*Dictionnary, error) {
	opts := badger.DefaultOptions(dir)
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)

	if err != nil {
		return nil, err
	}
	dict := &Dictionnary{
		db: db,
	}
	return dict, nil
}

func (d *Dictionnary) Close() {
	d.db.Close()
}
