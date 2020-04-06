package login

import (
	"database/sql"
	"fmt"
	"sivi/common"
	"sivi/entity"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func PostLogin(ListPgsql map[string]*sql.DB, c *gin.Context) (entity.ResponseHttp, error) {
	fmt.Println("postlogin ", "PostLogin")
	// Validate input
	var loginreq LoginRequest
	if err := c.BindJSON(&loginreq); err != nil {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_BADREQUEST,
			Message: common.ERRMSG_BADREQUEST,
			Data:    "",
		}, nil
	}

	if len(loginreq.UserName) == 0 {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_EMPTY_VALUE,
			Message: common.ERRMSG_EMPTY_VALUE,
			Data:    "",
		}, nil
	}

	if len(loginreq.Password) == 0 {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_EMPTY_VALUE,
			Message: common.ERRMSG_EMPTY_VALUE,
			Data:    "",
		}, nil
	}

	fmt.Println("loginreq UserName ", loginreq.UserName)
	fmt.Println("loginreq Password ", loginreq.Password)
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
						ua.ua_username = '%s' limit 1`, loginreq.UserName))
	defer results.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var resLogin LoginResponse
	if results.Next() {
		if err = results.Scan(&resLogin.ID,
			&resLogin.UserName,
			&resLogin.Password,
			&resLogin.IsStatus,
			&resLogin.CreateTime,
			&resLogin.CreatedBy,
			&resLogin.UpdateTime,
			&resLogin.UpdatedBy,
			&resLogin.UsertypeID); err != nil {
			panic(err.Error())
		}

		if resLogin.IsStatus == common.DBCODE_USER_STATUS_ACTIVE {
			fmt.Println("resLogin.Password ", resLogin.Password)
			fmt.Println("loginreq Password ", loginreq.Password)

			var scompare = strings.Compare(resLogin.Password, loginreq.Password)
			fmt.Println("scompare ", scompare)
			if scompare == 0 {
				resLogin.Password = "***************"
				return entity.ResponseHttp{
					Code:    common.ERRCODE_SUCCESS,
					Message: common.ERRMSG_SUCCESS,
					Data:    resLogin,
				}, nil
			} else {
				return entity.ResponseHttp{
					Code:    common.ERRCODE_WRONG_PASSWORD,
					Message: common.ERRMSG_WRONG_PASSWORD,
					Data:    "",
				}, nil
			}
		} else {
			return entity.ResponseHttp{
				Code:    common.ERRCODE_USERNAME_NOT_ACTIVE,
				Message: common.ERRMSG_USERNAME_NOT_ACTIVE,
				Data:    "",
			}, nil
		}
	} else {
		return entity.ResponseHttp{
			Code:    common.ERRCODE_NOTFOUND_IN_DB,
			Message: common.ERRMSG_NOTFOUND_IN_DB,
			Data:    "",
		}, nil
	}

	return entity.ResponseHttp{
		Code:    common.ERRCODE_UNDEFINED,
		Message: common.ERRMSG_UNDEFINED,
		Data:    "",
	}, nil
}
