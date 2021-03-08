package storage

import (
    "backend/model"
    bolt "go.etcd.io/bbolt"
    "golang.org/x/crypto/bcrypt"
    "strconv"
)

type ConfData struct {
    Key   string
    Value string
}

func (conf *ConfData) Init(s *Storage) error {
    passwdHash, err := bcrypt.GenerateFromPassword([]byte(model.DB_CONFIG_DEFAULT_PASSWD), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    defaultConf := []ConfData{
        {model.DB_CONFIG_PASSWD, string(passwdHash)},
        {model.DB_CONFIG_PORT, model.DB_CONFIG_DEFAULT_PORT},
    }

    err = s.DB.Update(func(tx *bolt.Tx) error {
        bucketConf, err := tx.CreateBucketIfNotExists([]byte(model.DB_NAME_CONF))
        if err != nil {
            return err
        }

        for _, config := range defaultConf {
            err = s.AddConfData(bucketConf, config.Key, config.Value)
            if err != nil {
                return err
            }
        }

        return nil
    })

    return err
}

func (conf *ConfData) Port() (int, error) {
    s, err := Open()
    if err != nil {
        return 0, err
    }
    defer s.Close()

    port := model.DB_CONFIG_DEFAULT_PORT

    err = s.DB.View(func(tx *bolt.Tx) error {
        bucketConf := tx.Bucket([]byte(model.DB_NAME_CONF))
        port = string(bucketConf.Get([]byte(model.DB_CONFIG_PORT)))
        return nil
    })
    if err != nil {
        return 0, err
    }

    return strconv.Atoi(port)
}

func (conf *ConfData) CheckPasswd(passwd string) (bool, error) {
    s, err := Open()
    if err != nil {
        return false, err
    }
    defer s.Close()

    var passwdHash string

    err = s.DB.View(func(tx *bolt.Tx) error {
        bucketConf := tx.Bucket([]byte(model.DB_NAME_CONF))
        passwdHash = string(bucketConf.Get([]byte(model.DB_CONFIG_PASSWD)))
        return nil
    })
    if err != nil {
        return false, err
    }

    if err = bcrypt.CompareHashAndPassword([]byte(passwdHash), []byte(passwd)); err != nil {
        if err.Error() == model.ERROR_PASSWD_WRONG {
            err = nil
        }
        return false, err
    } else {
        return true, nil
    }
}

func (conf *ConfData) UpdatePasswd(new string) error {
    s, err := Open()
    if err != nil {
        return err
    }
    defer s.Close()

    passwdHash, err := bcrypt.GenerateFromPassword([]byte(new), bcrypt.DefaultCost)

    err = s.DB.Update(func(tx *bolt.Tx) error {
        bucketConf := tx.Bucket([]byte(model.DB_NAME_CONF))
        err = bucketConf.Put([]byte(model.DB_CONFIG_PASSWD), passwdHash)
        return err
    })

    return nil
}
