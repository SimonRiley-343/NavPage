package storage

import (
	"backend/model"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB 		*sql.DB
}

func Open() (*Storage, error) {
	db, err := sql.Open("sqlite3", "./" + model.DB_FILE_NAME)
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

	row := s.DB.QueryRow(`SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?;`,
		model.DB_TABLE_USER)

	var isTableExist int
	if err = row.Scan(&isTableExist); err != nil {
		return err
	}

	if isTableExist == 0 {
		_, err = s.DB.Exec(`CREATE TABLE user (id INTEGER NOT NULL PRIMARY KEY, user TEXT, passwd TEXT);`)
		if err != nil {
			return err
		}
		ud := UserData{}
		if err = ud.Init(); err != nil {
			return err
		}
	}

	return nil
}


