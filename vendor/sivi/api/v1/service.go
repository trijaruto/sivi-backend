package v1

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sivi/api/v1/login"
	"sivi/api/v1/signup"
	"sivi/common"
	"strings"
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

func ValidateTokenHeader(c *gin.Context) {

	var sivitoken = strings.Replace(c.Request.Header.Get("Token"), "Bearer ", "", 1)
	fmt.Println("token : ", sivitoken)

	token, err := jwt.Parse(sivitoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("token.key")), nil
	})
	if err != nil {
		panic(err.Error())
	}

	// // When using `Parse`, the result `Claims` would be a map.
	fmt.Println("token.Claims : ", token.Claims)

	//TOKEN REQUEST
	reqTok := entity.ResponseToken{}
	token, err = jwt.ParseWithClaims(sivitoken, &reqTok, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("token.key")), nil
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("reqTok.Code : ", reqTok.Code)
	fmt.Println("reqTok.Message : ", reqTok.Message)
	fmt.Println("reqTok.Data : ", reqTok.Data)

	var timeNow = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("timeNow date : ", time.Now())
	fmt.Println("timeNow format string : ", timeNow)
	fmt.Println("expired format string : ", reqTok.Expired)

	datetimenow, errdatetimenow := time.Parse("2006-01-02 15:04:05", timeNow)
	if errdatetimenow != nil {
		panic(errdatetimenow.Error())
	}
	fmt.Println("datetimenow date : ", datetimenow)
	datetimeexpired, errdatetimeexpired := time.Parse("2006-01-02 15:04:05", reqTok.Expired)
	if errdatetimeexpired != nil {
		panic(errdatetimeexpired.Error())
	}
	fmt.Println("datetimeexpired date : ", datetimeexpired)
	if datetimenow.Before(datetimeexpired) {
		fmt.Println("datetimeexpired false ")
		fmt.Println(datetimeexpired, "Before : ", datetimenow)

		//reqDataLoginToken := reqTok.Data.(map[string]interface{})
		// for key, value := range reqDataLoginToken {
		// 	if key == "id" {
		// 		req.UserId = int(value.(float64))
		// 		fmt.Println("req.UserId : ", req.UserId)
		// 	}
		//}
	}
}
