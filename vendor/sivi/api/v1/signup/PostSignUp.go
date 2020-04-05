package signup

import (
	"database/sql"
	"fmt"
	"sivi/common"

	"github.com/gin-gonic/gin"
	"gitlab.smartfren.com/common/rest-client/entity"
)

func PostSignUp(ListPgsql map[string]*sql.DB, c *gin.Context) (entity.ResponseHttp, error) {
	fmt.Println("PostSignUp ", "PostSignUp")
	return entity.ResponseHttp{
		Code:    common.ERRCODE_UNDEFINED,
		Message: common.ERRMSG_UNDEFINED,
		Data:    "PostSignUp",
	}, nil
}
