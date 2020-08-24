package storage

import (
	"backend/model"
	"log"

	"github.com/asdine/storm"
)

type Client struct {
	db *storm.DB
}

func Open() (*Client, error) {
	var db *storm.DB
		db, err := storm.Open(model.DB_FILE_NAME)
	if err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}

func (c *Client) Close() {
	if err := c.db.Close(); err != nil {
		log.Println(err)
	}
}
