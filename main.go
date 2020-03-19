package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
		logInfo.Println("$PORT must be set")
	} else {
		logInfo.Println("port : ", port)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
