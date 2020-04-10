package signup

import (
	"database/sql"
	"fmt"
	"log"
	"sivi/common"
	"sivi/entity"
	"sivi/security"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func PostSignUp(ListPgsql map[string]*sql.DB, c *gin.Context) (entity.ResponseHttp, error) {
	fmt.Println("PostSignUp ", "PostSignUp")

	var signupreq SignUpRequest
	if err := c.BindJSON(&signupreq); err != nil {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_BADREQUEST,
			Message: common.ERRMSG_BADREQUEST,
			Data:    "",
		}, nil
	}

	if len(signupreq.UserName) == 0 {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_EMPTY_VALUE,
			Message: common.ERRMSG_EMPTY_VALUE,
			Data:    "",
		}, nil
	}

	if len(signupreq.Password) == 0 {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_EMPTY_VALUE,
			Message: common.ERRMSG_EMPTY_VALUE,
			Data:    "",
		}, nil
	}

	if len(signupreq.Password) < common.APP_MIN_PASSWORD_LENGTH {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_MIN_PASSWORD_LENGTH,
			Message: common.ERRMSG_MIN_PASSWORD_LENGTH,
			Data:    "",
		}, nil
	}

	fmt.Println("loginreq UserName ", signupreq.UserName)
	fmt.Println("loginreq Password ", signupreq.Password)
	results, err := ListPgsql[viper.GetString("database.heroku.postgresql.name")].Query(
		fmt.Sprintf(`select 	
						ua.ua_id, 
						ua.ua_username,	
						ua_userpassword,					
						ua.ua_isstatus,
						ua_createtime,
						ua_createdby,
						ua_updatetime,
						ua_updatedby,
						ua_usertype_id				
					from user_account as ua					
					where 						
						ua.ua_username = '%s' limit 1`, signupreq.UserName))
	defer results.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_ALREADYEXISTS,
			Message: common.ERRMSG_ALREADYEXISTS,
			Data:    "",
		}, nil
	} else {
		return InsertUserAccount(ListPgsql, signupreq)
	}

	return entity.ResponseHttp{
		Code:    common.ERRCODE_UNDEFINED,
		Message: common.ERRMSG_UNDEFINED,
		Data:    "PostSignUp",
	}, nil
}

func InsertUserAccount(ListPgsql map[string]*sql.DB, sReq SignUpRequest) (entity.ResponseHttp, error) {
	fmt.Println("loginreq UserName ", sReq.UserName)
	fmt.Println("loginreq Password ", sReq.Password)
	fmt.Println("time.Now() ", time.Now())

	db, err := ListPgsql[viper.GetString("database.heroku.postgresql.name")].Begin()
	if err != nil {
		panic(err.Error())
	}

	resUa, errUa := db.Prepare("INSERT INTO user_account (ua_username, ua_userpassword, ua_isstatus, ua_createtime, ua_createdby, ua_updatetime, ua_updatedby, ua_usertype_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING ua_id;")
	if errUa != nil {
		db.Rollback()
		panic(errUa.Error())
	}
	defer resUa.Close()

	var uaID int
	err = resUa.QueryRow(
		sReq.UserName,
		security.GeneratePassword(fmt.Sprintf("%s:%s", sReq.UserName, sReq.Password)),
		common.DBCODE_USER_STATUS_NOT_ACTIVE,
		time.Now(),
		sReq.UserName,
		time.Now(),
		sReq.UserName,
		common.DBCODE_USER_TYPE_USER,
	).Scan(&uaID)
	if err != nil {
		db.Rollback()
		log.Fatal(err)
	}

	db.Commit()

	return entity.ResponseHttp{
		Code:    common.ERRCODE_SUCCESS,
		Message: common.ERRMSG_SUCCESS,
		Data:    fmt.Sprintf("{ua_id : '%s'}", strconv.Itoa(uaID)),
	}, nil
}
