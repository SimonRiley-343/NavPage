package storage

import (
    "backend/model"
    "errors"
    bolt "go.etcd.io/bbolt"
    "log"
    "time"
)

type Storage struct {
    DB *bolt.DB
}

func Open() (*Storage, error) {
    db, err := bolt.Open(model.DB_FILE_NAME, 0600, &bolt.Options{Timeout: 3 * time.Second})
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

    bucketList := []string{model.DB_NAME_CONF, model.DB_NAME_PAGE}

    for _, bucketName := range bucketList {
        isBucketExist := false
        err = s.DB.View(func(tx *bolt.Tx) error {
            if tx.Bucket([]byte(bucketName)) != nil {
                isBucketExist = true
            }
            return nil
        })

        if err != nil {
            return err
        }

        if !isBucketExist {
            if err = createBucket(s, bucketName); err != nil {
                return err
            }
        }
    }

    return nil
}

func createBucket(s *Storage, bucketName string) error {
    err := s.DB.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(bucketName))
        return err
    })

    if err != nil {
        return err
    }

    switch bucketName {
    case model.DB_NAME_CONF:
        confData := ConfData{}
        err = confData.Init(s)
    case model.DB_NAME_PAGE:
        pageData := PageData{}
        err = pageData.Init(s)
    default:
        err = errors.New("Unknown bucket name")
    }
    if err != nil {
        return err
    }

    return err
}
