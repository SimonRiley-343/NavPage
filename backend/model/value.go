package model

const (
	DB_FILE_NAME           = "nav_page.db"
	DB_TABLE_USER          = "user"
	DB_USER_DEFAULT_ID     = 1001
	DB_USER_DEFAULT_USER   = "admin"
	DB_USER_DEFAULT_PASSWD = "passwd"

	POST_FORM_USER   = "user"
	POST_FORM_PASSWD = "passwd"

	REQ_MSG_PARAM_ERR     = "Param error"
	REQ_MSG_LOGIN_SUCCESS = "Success login"
	REQ_MSG_LOGIN_FAILED  = "Failed login"
	REQ_MSG_UNKNOWN_ERR   = "Unknown error"

	REQ_CODE_SUCCESS      = 0
	REQ_CODE_PARAM_ERR    = 101
	REQ_CODE_LOGIN_FAILED = 102
	REQ_CODE_UNKNOWN_ERR  = 201

	ERROR_PASSWD_WRONG = "crypto/bcrypt: hashedPassword is not the hash of the given password"
)
