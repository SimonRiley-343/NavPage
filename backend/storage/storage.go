package storage

import (
	"backend/model"
	"log"

	"github.com/asdine/storm"
)

type Client struct {
	DB *storm.DB
}

func Open() (*Client, error) {
	var db *storm.DB
		db, err := storm.Open(model.DB_FILE_NAME)
	if err != nil {
		return nil, err
	}
	return &Client{DB: db}, nil
}

func (c *Client) Close() {
	if err := c.DB.Close(); err != nil {
		log.Println(err)
	}
}
