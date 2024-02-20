package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DatabaseInstance *sqlx.DB
	DatabaseErr      = make(chan error)
)

func ConnectInstance() (err error) {
	DatabaseInstance, err = sqlx.Open("sqlite3", "/data/rinha_api.sqlite?_journal=WAL&_timeout=5000&_fk=true")
	if err != nil {
		return
	}
	if err = DatabaseInstance.Ping(); err != nil {
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

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			slice_content := strings.Split(string(content), "\n--")
			for i := range slice_content {
				if _, err = DatabaseInstance.Exec(slice_content[i]); err != nil {
					panic(err)
				}

			}

		}
	}

	return
}
