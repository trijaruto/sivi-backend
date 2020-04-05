package login

import (
	"database/sql"
	"fmt"
	"sivi/common"

	"github.com/gin-gonic/gin"
	"gitlab.smartfren.com/common/rest-client/entity"
)

func PostLogin(ListPgsql map[string]*sql.DB, c *gin.Context) (entity.ResponseHttp, error) {
	fmt.Println("PostLogin ", "PostLogin")
	return entity.ResponseHttp{
		Code:    common.ERRCODE_UNDEFINED,
		Message: common.ERRMSG_UNDEFINED,
		Data:    "PostLogin",
	}, nil
}
