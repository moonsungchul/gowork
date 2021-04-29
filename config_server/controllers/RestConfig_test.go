package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/moonsungchul/configserver/controllers"
	"github.com/moonsungchul/configserver/models"
)

func setupRouter() *gin.Engine {
	store := models.MysqlStore{Host: "172.17.0.3", Port: 3306,
		Dbname: "fms_config", User: "moonstar", Passwd: "wooag01"}
	//db := store.Open("172.17.0.3", 3306, "fms_config", "moonstar", "wooag01")
	db := store.OpenDB()
	store.CreateTable(db)
	r_conf := controllers.RestConfig{&store}
	r_common := controllers.CommonInfo{}
	rr := gin.Default()
	rr.GET("/api/version", r_common.GetVersion)
	rr.POST("/api/config", r_conf.AddConfig)
	rr.GET("/api/config/:cname/:section/:key", r_conf.GetConfig)
	return rr
}

func TestVetsion(t *testing.T) {
	rr := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/version", nil)
	rr.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	fmt.Println("test unit test ")

	jj := controllers.JConfig{ConfigName: "Test", Section: "TestSection", Key: "Key1", Value: "Value1"}
	jb, _ := json.Marshal(jj)
	w1 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/api/config", bytes.NewReader(jb))

	rr.ServeHTTP(w1, req2)
	fmt.Println(w1.Body.String())
	fmt.Println("test unit test ")
}

func TestGetConfigValue(t *testing.T) {
	rr := setupRouter()

	w1 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/config/Test/TestSection/Key1", nil)

	rr.ServeHTTP(w1, req2)
	fmt.Println(">>>>> :", w1.Body.String())
	fmt.Println("test unit test ")
}
