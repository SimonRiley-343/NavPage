package model

const (
	DB_FILE_NAME             = "nav_page.db"
	DB_PAGE_DEFAULT_ID       = 1001
	DB_PAGE_DEFAULT_NAME     = "Example"
	DB_PAGE_DEFAULT_DESC     = "Example description"
	DB_PAGE_DEFAULT_URL      = "/"
	DB_CONFIG_PASSWD         = "passwd"
	DB_CONFIG_PORT           = "port"
	DB_CONFIG_DEFAULT_PASSWD = "passwd"
	DB_CONFIG_DEFAULT_PORT   = "8080"

	POST_FORM_PASSWD = "passwd"

	REQ_MSG_PARAM_ERR    = "Param error"
	REQ_MSG_LOGIN_FAILED = "Failed login"
	REQ_MSG_PAGE_NO      = "No page"
	REQ_MSG_SUCCESS      = "Success"
	REQ_MSG_UNKNOWN_ERR  = "Unknown error"

	REQ_CODE_SUCCESS      = 0
	REQ_CODE_PARAM_ERR    = 101
	REQ_CODE_LOGIN_FAILED = 102
	REQ_CODE_NO_PAGE      = 201
	REQ_CODE_UNKNOWN_ERR  = 201

	ERROR_PASSWD_WRONG = "crypto/bcrypt: hashedPassword is not the hash of the given password"
)
