package common

const (
	ERRCODE_UNDEFINED     int8 = -1
	ERRCODE_SUCCESS       int8 = 0
	ERRCODE_NOTFOUND      int8 = 1
	ERRCODE_ALREADYEXISTS int8 = 2
	ERRCODE_BADROUTING    int8 = 3
)

const (
	ERRMSG_UNDEFINED     string = "undefined error"
	ERRMSG_SUCCESS       string = "success"
	ERRMSG_NOTFOUND      string = "not found"
	ERRMSG_ALREADYEXISTS string = "already exists"
	ERRMSG_BADROUTING    string = "inconsistent mapping between route and handler"
)

const (
	DBCODE_USER_STATUS_UNDEFINED  int = -1
	DBCODE_USER_STATUS_NOT_ACTIVE int = 0
	DBCODE_USER_STATUS_ACTIVE     int = 1
)
