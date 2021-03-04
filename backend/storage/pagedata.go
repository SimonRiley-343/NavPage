package storage

import (
	"backend/model"
	"encoding/json"
	bolt "go.etcd.io/bbolt"
)

type PageData struct {
    name string
    desc string
    url  string
    cat  string
}

func (pd *PageData) Init(s *Storage) error {
    err := s.DB.Update(func(tx *bolt.Tx) error {
        bucketPage, err := tx.CreateBucketIfNotExists([]byte(model.DB_NAME_PAGE))
        if err != nil {
            return err
        }

        err = s.AddPageData(bucketPage, model.DB_PAGE_DEFAULT_NAME, model.DB_PAGE_DEFAULT_CAT,
            model.DB_PAGE_DEFAULT_DESC, model.DB_PAGE_DEFAULT_URL)
        return err
    })

    return err
}

func (pd *PageData) GetAllPage() ([]model.Pages, error) {
    s, err := Open()
    if err != nil {
        return nil, err
    }
    defer s.Close()

    var pageList []model.Pages

    err = s.DB.View(func(tx *bolt.Tx) error {
        bucketPage := tx.Bucket([]byte(model.DB_NAME_PAGE))
        err = bucketPage.ForEach(func(pageIdB, pageDataB []byte) error {
            var pageData map[string]interface{}
            err := json.Unmarshal(pageDataB, &pageData)
            if err != nil {
                return err
            }
            pageList = append(pageList, model.Pages{
                Id: pageData["id"].(string),
                Name: pageData["name"].(string),
                Cat: pageData["cat"].(string),
                Desc: pageData["desc"].(string),
                Url: pageData["url"].(string),
            })
            return nil
        })
        return err
    })

    if err != nil {
        return nil, err
    }

    return pageList, nil
}
