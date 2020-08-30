package storage

import (
	"backend/model"
	"log"
)

type PageData struct {
	id   int
	name string
	desc string
	url  string
}

func (pd *PageData) Init() error {
	s, err := Open()
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = s.DB.Exec(`DELETE FROM page;`)
	if err != nil {
		return err
	}

	_, err = s.DB.Exec(`INSERT INTO page (id, name, desc, url) VALUES (?, ?, ?, ?);`,
		model.DB_PAGE_DEFAULT_ID, model.DB_PAGE_DEFAULT_NAME, model.DB_PAGE_DEFAULT_DESC, model.DB_PAGE_DEFAULT_URL)
	if err != nil {
		return err
	}
	return nil
}

func (pd *PageData) GetAllPage() ([]model.Pages, error) {
	s, err := Open()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	rows, err := s.DB.Query(`SELECT id, name, desc, url FROM page;`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	var pageList []model.Pages
	var id int
	var name, desc, url string

	for rows.Next() {
		if err = rows.Scan(&id, &name, &desc, &url); err != nil {
			return nil, err
		}
		pageList = append(pageList, model.Pages{Id: id, Name: name, Desc: desc, Url: url})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return pageList, nil
}