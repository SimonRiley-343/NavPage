package storage

import (
	"backend/model"
	"golang.org/x/crypto/bcrypt"
)

type ConfData struct {
	Name  string
	Value string
}

func (conf *ConfData) Init() error {
	s, err := Open()
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = s.DB.Exec(`DELETE FROM config;`)
	if err != nil {
		return err
	}

	passwdHash, err := bcrypt.GenerateFromPassword([]byte(model.DB_CONFIG_DEFAULT_PASSWD), bcrypt.DefaultCost)

	defaultConf := []ConfData{
		{model.DB_CONFIG_PASSWD, string(passwdHash)},
		{model.DB_CONFIG_PORT, model.DB_CONFIG_DEFAULT_PORT},
	}

	for _, config := range defaultConf {
		_, err = s.DB.Exec(`INSERT INTO config (name, value) VALUES (?, ?);`,
			config.Name, config.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (conf *ConfData) Port() (int, error) {
	s, err := Open()
	if err != nil {
		return 0, err
	}
	defer s.Close()

	var port int
	rows, err := s.DB.Query(`SELECT value FROM config WHERE name = ?;`,
		model.DB_CONFIG_PORT)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&port); err != nil {
			return 0, err
		}
	}

	if rows.Err() != nil {
		return 0, rows.Err()
	}

	return port, nil
}

func (conf *ConfData) CheckPasswd(passwd string) (bool, error) {
	s, err := Open()
	if err != nil {
		return false, err
	}
	defer s.Close()

	var passwdHash string
	rows, err := s.DB.Query(`SELECT value FROM config WHERE name = ?;`,
		model.DB_CONFIG_PASSWD)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&passwdHash); err != nil {
			return false, err
		}
	}

	if rows.Err() != nil {
		return false, rows.Err()
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
	if _, err = s.DB.Exec(`UPDATE config SET value = ? WHERE name = ?;`,
		string(passwdHash), model.DB_CONFIG_PASSWD); err != nil {
		return err
	}
	return nil
}
