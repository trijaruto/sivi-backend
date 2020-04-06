package login

import (
	"database/sql"
	"fmt"
	"sivi/common"
	"sivi/entity"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func PostLogin(ListPgsql map[string]*sql.DB, c *gin.Context) (entity.ResponseHttp, error) {
	fmt.Println("postlogin ", "PostLogin")
	results, err := ListPgsql[viper.GetString("database.heroku.postgresql.name")].Query(
		fmt.Sprintf(`select 	
						ua.ua_id, 
						ua.ua_username,						
						ua.ua_isstatus				
					from user_account as ua					
					where 						
						ua.ua_isstatus = '1' limit 1`))
	defer results.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var resLogin LoginResponse
	for results.Next() {
		if err = results.Scan(&resLogin.ID, &resLogin.UserName, &resLogin.IsStatus); err != nil {
			panic(err.Error())
		}

		if resLogin.IsStatus == common.DBCODE_USER_STATUS_ACTIVE {
			return entity.ResponseHttp{
				Code:    common.ERRCODE_SUCCESS,
				Message: common.ERRMSG_SUCCESS,
				Data:    resLogin,
			}, nil
		}

	}

	return entity.ResponseHttp{
		Code:    common.ERRCODE_NOTFOUND_IN_DB,
		Message: common.ERRMSG_NOTFOUND_IN_DB,
		Data:    "",
	}, nil
}
