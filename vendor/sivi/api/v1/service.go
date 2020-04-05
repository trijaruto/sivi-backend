package v1

import (
	"database/sql"
	"log"
	"net/http"
	"sivi/api/v1/login"
	"sivi/api/v1/signup"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusOK, result)
	}
}

func (s *ServiceStruct) PostSignUp(c *gin.Context) {
	result, rerror := signup.PostSignUp(s.ListPgsql, c)
	if rerror != nil {
		panic(rerror.Error())
	} else {
		c.JSON(http.StatusOK, result)
	}
}
