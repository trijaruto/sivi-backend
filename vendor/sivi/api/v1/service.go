package v1

import (
	"database/sql"
	"log"
	"net/http"
	"sivi/api/v1/login"
	"sivi/api/v1/signup"
	"sivi/common"
	"time"

	"sivi/entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ServiceStruct struct {
	ListPgsql map[string]*sql.DB
	LogInfo   *log.Logger
}

func (s *ServiceStruct) PostLoginService(c *gin.Context) {
	result, rerror := login.PostLogin(s.ListPgsql, c)
	if rerror != nil {
		panic(rerror.Error())
	} else {
		LoginTokenHeader(c, result)
		c.JSON(http.StatusOK, result)
	}
}

func (s *ServiceStruct) PostSignUpService(c *gin.Context) {
	result, rerror := signup.PostSignUp(s.ListPgsql, c)
	if rerror != nil {
		panic(rerror.Error())
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func LoginTokenHeader(c *gin.Context, res entity.ResponseHttp) {
	second := time.Duration(viper.GetInt("token.expired.second")) * time.Second
	var timeExpired = time.Now().Add(second).Format("2006-01-02 15:04:05")
	if res.Code == common.ERRCODE_SUCCESS {
		// HEADER AND PAYLOAD
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &entity.ResponseToken{
			Code:    common.ERRCODE_SUCCESS,
			Message: common.ERRMSG_SUCCESS,
			Expired: timeExpired,
			Data:    res.Data,
		})
		// SIGNATURE
		tokenstring, errtoken := token.SignedString([]byte(viper.GetString("token.key")))
		if errtoken != nil {
			panic(errtoken.Error())
		}
		c.Header("Token", "Bearer "+tokenstring)
	}

}
