package common

const (
	ERRCODE_UNDEFINED           int8 = -1
	ERRCODE_SUCCESS             int8 = 0
	ERRCODE_NOTFOUND            int8 = 1
	ERRCODE_ALREADYEXISTS       int8 = 2
	ERRCODE_BADREQUEST          int8 = 3
	ERRCODE_NOTFOUND_IN_DB      int8 = 4
	ERRCODE_USERNAME_NOT_ACTIVE int8 = 5
	ERRCODE_WRONG_PASSWORD      int8 = 6
	ERRCODE_EMPTY_VALUE         int8 = 7
)

const (
	ERRMSG_UNDEFINED           string = "undefined error"
	ERRMSG_SUCCESS             string = "success"
	ERRMSG_NOTFOUND            string = "not found"
	ERRMSG_ALREADYEXISTS       string = "already exists"
	ERRMSG_BADREQUEST          string = "bad request"
	ERRMSG_NOTFOUND_IN_DB      string = "not found in db"
	ERRMSG_USERNAME_NOT_ACTIVE string = "username not active"
	ERRMSG_WRONG_PASSWORD      string = "wrong password"
	ERRMSG_EMPTY_VALUE         string = "empty value"
)

const (
	DBCODE_USER_STATUS_UNDEFINED  int = -1
	DBCODE_USER_STATUS_NOT_ACTIVE int = 0
	DBCODE_USER_STATUS_ACTIVE     int = 1
)
