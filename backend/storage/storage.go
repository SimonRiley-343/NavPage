package storage

import (
	"backend/model"
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func Open() (*Storage, error) {
	db, err := sql.Open("sqlite3", "./"+model.DB_FILE_NAME)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Close() {
	if err := s.DB.Close(); err != nil {
		log.Println(err)
	}
}

func CheckTable() error {
	s, err := Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer s.Close()

	tableList := []string{"config", "page"}

	for _, table := range tableList {
		row := s.DB.QueryRow(`SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?;`,
			table)

		var isTableExist int
		if err = row.Scan(&isTableExist); err != nil {
			return err
		}

		if isTableExist == 0 {
			if err = createTable(table, s); err != nil {
				return err
			}
		}
	}
	return nil
}

func createTable(table string, s *Storage) error {
	switch table {
	case "page":
		_, err := s.DB.Exec(`CREATE TABLE page (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, desc TEXT, url TEXT NOT NULL, cat TEXT);`)
		if err != nil {
			return err
		}
		pd := PageData{}
		if err = pd.Init(); err != nil {
			return err
		}
	case "config":
		_, err := s.DB.Exec(`CREATE TABLE config (name TEXT NOT NULL PRIMARY KEY, value TEXT NOT NULL);`)
		if err != nil {
			return err
		}
		conf := ConfData{}
		if err = conf.Init(); err != nil {
			return err
		}
	default:
		return errors.New("unknown table name")
	}
	return nil
}
