package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	_ = viper.ReadInConfig()
}

type ResponseObj struct {
	Operation string      `json:"operation"`
	Err       int         `json:"error"`
	Timestamp time.Time   `json:"timestamp"`
	Path      string      `json:"path"`
	Content   interface{} `json:"content"`
}

func main() {
	router := httprouter.New()

	router.GET("/*path", readHandler)

	// File watcher
	router.GET("/filewatch", fileWatchHandler)
	router.PUT("/*path", writeHandler)
	router.DELETE("/*path", deleteHandler)
	router.POST("/*path", modifyHandler)

	log.Println("Summer server listening at port" + ":" + viper.Get("appPort").(string))

	log.Fatal(http.ListenAndServe(":"+viper.Get("appPort").(string), middleware(router)))
}
