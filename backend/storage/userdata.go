package storage

import (
	"backend/model"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	id 		int
	user 	string
	passwd 	string
}

func (ud *UserData) Init() error {
	s, err := Open()
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = s.DB.Exec(`DELETE FROM user;`)
	if err != nil {
		return err
	}

	passwdHash, err := bcrypt.GenerateFromPassword([]byte(model.DB_USER_DEFAULT_PASSWD), bcrypt.DefaultCost)

	_, err = s.DB.Exec(`INSERT INTO user (id, user, passwd) VALUES (?, ?, ?);`,
		model.DB_USER_DEFAULT_ID, model.DB_USER_DEFAULT_USER, string(passwdHash))
	if err != nil {
		return err
	}
	return nil
}

func (ud *UserData) CheckPasswd(id int, passwd string) (bool, error) {
	s, err := Open()
	if err != nil {
		return false, err
	}
	defer s.Close()

	var passwdHash string
	rows, err := s.DB.Query(`SELECT passwd FROM user WHERE id = ?;`,
		id)
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

func (ud *UserData) UpdateUser(id int, new string) error {
	s, err := Open()
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = s.DB.Exec(`UPDATE user SET user = ? WHERE id = ?;`,
		new, id)
	if err != nil {
		return err
	}
	return nil
}
