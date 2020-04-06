package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	v1 "sivi/api/v1"
	"sivi/connection"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/spf13/viper"
)

func main() {

	fmt.Println("Hello World")
	//1. Flag Parse
	fmt.Println("1. Flag Parse")
	var (
		profile    = flag.String("profile", "dev", "Profile app environment")
		configPath = flag.String("configpath", "config/", "Config file path")
		//logOutput  = flag.String("logoutput", "file", "Output log")
	)
	flag.Parse()

	//2. Load Config With Viper
	fmt.Println("2. Load Config With Viper")
	viper.SetConfigType("yaml")
	var configFileName []string
	configFileName = append(configFileName, "config-")
	configFileName = append(configFileName, *profile)
	viper.SetConfigName(strings.Join(configFileName, ""))
	viper.AddConfigPath(*configPath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	} else {
		fmt.Println("   configfile : ")
		fmt.Println("   	configFileName : ", configFileName)
	}

	//3. Logging init
	fmt.Println("3. Logging Init")
	var logFileName []string
	//currentDate := time.Now()
	logFileName = append(logFileName, viper.GetString("log.path"))
	//logFileName = append(logFileName, currentDate.Format("2006-01-02"))
	//logFileName = append(logFileName, "/")
	os.MkdirAll(strings.Join(logFileName, ""), os.ModePerm)
	logFileName = append(logFileName, viper.GetString("log.filename"))
	logfile, err := os.OpenFile(strings.Join(logFileName, ""), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		fmt.Println("   log : ")
		fmt.Println("   	logFileName : ", logFileName)
	}
	defer logfile.Close()

	logInfo := log.New(logfile, "loginfo : ", log.LstdFlags)
	logInfo.Println("configFileName : ", configFileName)
	logInfo.Println("logfilename : ", logFileName)

	//4. Connection DB Localhost
	// fmt.Println("4. Connection DB PostgreSQL Localhost")
	// dbPgsql := connection.Pgsql{}
	// dbPgsql.PgsqlMultipleConnection(connection.PgsqlParams{
	// 	Name: viper.GetString("database.heroku.postgresql.name"), Database: viper.GetString("database.localhost.postgresql.database"), Host: viper.GetString("database.localhost.postgresql.hostname"), Port: viper.GetInt("database.localhost.postgresql.port"), User: viper.GetString("database.localhost.postgresql.username"), Password: viper.GetString("database.localhost.postgresql.password"), Schema: viper.GetString("database.localhost.postgresql.schema"), Driver: viper.GetString("database.localhost.postgresql.driver"),
	// })
	// err = dbPgsql.ListPgsql[viper.GetString("database.heroku.postgresql.name")].Ping()

	// logInfo.Println("db : ", viper.GetString("database.localhost.postgresql.driver"))
	// logInfo.Println("Name : ", viper.GetString("database.heroku.postgresql.name"))

	//4. Connection DB HEROKU
	fmt.Println("4. Connection DB PostgreSQL Heroku")
	dbPgsql := connection.Pgsql{}
	dbPgsql.HerokuPgsqlMultipleConnection(connection.PgsqlParams{
		Name: viper.GetString("database.heroku.postgresql.name"), Database: viper.GetString("database.heroku.postgresql.database"), Host: viper.GetString("database.heroku.postgresql.hostname"), Port: viper.GetInt("database.heroku.postgresql.port"), User: viper.GetString("database.heroku.postgresql.username"), Password: viper.GetString("database.heroku.postgresql.password"), Schema: viper.GetString("database.heroku.postgresql.schema"), Driver: viper.GetString("database.heroku.postgresql.driver"), URI: viper.GetString("database.heroku.postgresql.uri"),
	})
	err = dbPgsql.ListPgsql[viper.GetString("database.heroku.postgresql.name")].Ping()

	logInfo.Println("db : ", viper.GetString("database.heroku.postgresql.driver"))
	logInfo.Println("Name : ", viper.GetString("database.heroku.postgresql.name"))

	if err != nil {
		logInfo.Println("connectionDB", "failed")
		fmt.Println("connectionDB", "failed")
		panic(err)
	} else {
		logInfo.Println("connectionDB", "Success")
		fmt.Println("connectionDB", "Success")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
		logInfo.Println("$PORT must be set")
	} else {
		logInfo.Println("App running on port : ", port)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"sivi":      "!1qaz)0plm8",
		"trijaruto": "asdlkjmnbzxc",
	}))

	if viper.GetString("api.version") == "v1" {
		apiStruct := &v1.ServiceStruct{ListPgsql: dbPgsql.ListPgsql, LogInfo: logInfo}
		rv1 := authorized.Group(fmt.Sprintf("/%s", viper.GetString("api.version")))
		{
			rv1.POST(fmt.Sprintf("/%s", viper.GetString("api.service.apilogin.path")), apiStruct.PostLoginService)
			rv1.POST(fmt.Sprintf("/%s", viper.GetString("api.service.apisignup.path")), apiStruct.PostSignUp)
		}
	} else {
		fmt.Println("API Version ", "no version")
	}

	router.Run(":" + port)
}
