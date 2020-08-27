package model

const (
	DB_FILE_NAME         	= "nav_page.db"
	DB_TABLE_USER        	= "user"
	DB_USER_DEFAULT_ID	 	= 1001
	DB_USER_DEFAULT_USER 	= "admin"
	DB_USER_DEFAULT_PASSWD 	= "passwd"

	ERROR_PASSWD_WRONG		= "crypto/bcrypt: hashedPassword is not the hash of the given password"
)
