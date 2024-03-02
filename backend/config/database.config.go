package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ReaderDB *sqlx.DB
	WriterDB *sqlx.DB

	ErrWriterDB = make(chan error)
	ErrReaderDB = make(chan error)
)

func ConnectInstance() (connection *sqlx.DB, err error) {
	connection, err = sqlx.Open("sqlite3", "/data/rinha_api.sqlite?_journal=WAL&_timeout=5000&_fk=true")
	if err != nil {
		return
	}
	if err = connection.Ping(); err != nil {
		return
	}

	return
}

func ExecMigration() (err error) {
	dir := "/data/schemas/"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	WriterDB.Exec(`
	PRAGMA automatic_index = ON;
	PRAGMA cache_size = 32768;
	PRAGMA cache_spill = OFF;
	PRAGMA foreign_keys = ON;
	PRAGMA journal_size_limit = 67110000;
	PRAGMA locking_mode = NORMAL;
	PRAGMA page_size = 4096;
	PRAGMA recursive_triggers = ON;
	PRAGMA secure_delete = ON;
	PRAGMA synchronous = NORMAL;
	PRAGMA temp_store = MEMORY;
	PRAGMA journal_mode = WAL;
	PRAGMA wal_autocheckpoint = 16384;
	`)

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			slice_content := strings.Split(string(content), "\n--")
			for i := range slice_content {
				if _, err = WriterDB.Exec(slice_content[i]); err != nil {
					panic(err)
				}

			}

		}
	}

	return
}
