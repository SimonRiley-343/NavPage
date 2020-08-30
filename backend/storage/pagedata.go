package storage

import (
	"backend/model"
	"log"
)

type PageData struct {
	id   int
	name string
	url  string
	img  string
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

	_, err = s.DB.Exec(`INSERT INTO page (id, name, url, img) VALUES (?, ?, ?, ?);`,
		model.DB_PAGE_DEFAULT_ID, model.DB_PAGE_DEFAULT_NAME, model.DB_PAGE_DEFAULT_URL, model.DB_PAGE_DEFAULT_IMG)
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

	rows, err := s.DB.Query(`SELECT id, name, url, img FROM page;`)
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
	var name, url, img string

	for rows.Next() {
		if err = rows.Scan(&id, &name, &url, &img); err != nil {
			return nil, err
		}
		pageList = append(pageList, model.Pages{Id: id, Name: name, Url: url, Img: img})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return pageList, nil
}
