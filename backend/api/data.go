package api

import (
	"backend/model"
	"backend/storage"
	"errors"
	"log"
)

type Data struct {
	Client *storage.Client
}

func (d *Data) CheckData() error {
	var err error
	d.Client, err = storage.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Client.Close()

	bucketIsExist := false

	for _, bucket := range d.Client.DB.Bucket() {
		if bucket == model.DB_BUCKET_NAME {
			bucketIsExist = true
		}
	}

	if !bucketIsExist {
		//TODO: Init bucket
		return errors.New("bucket " + model.DB_BUCKET_NAME + " not exist")
	}

	//TODO: Replace with map, inlcude key name and default value
	keySlice := []string{
		model.DB_KEY_USER,
		model.DB_KEY_PASSWD,
		model.DB_KEY_TOKEN,
		model.DB_KEY_PAGES,
	}

	for _, key := range keySlice {
		exist, _ := d.checkConf(model.DB_BUCKET_NAME, key, nil)
		if !exist {
			return errors.New("key " + key + " not exist")
		}
	}
	return nil
}

func (d *Data) checkConf(bucket string, key string, defaultValue interface{}) (bool, error) {
	//TODO: Create key then set default value by input
	exist, err := d.Client.DB.KeyExists(bucket, key)
	return exist, err
}
