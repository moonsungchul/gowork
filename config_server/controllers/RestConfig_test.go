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
	"github.com/stretchr/testify/assert"
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
	rr.DELETE("/api/config/:cname/:section/:key", r_conf.DeleteProperty)
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

	var cc models.ConfigProperty
	if err := json.Unmarshal(w1.Body.Bytes(), &cc); err != nil {
		panic(err)
	}
	fmt.Println("Property : ", cc)
	assert.Equal(t, cc.SKey, "Key1", "기대값과 결과값이 다릅니다. ")
}

func TestDeleteConfigProperty(t *testing.T) {
	rr := setupRouter()
	jj := controllers.JConfig{ConfigName: "Test", Section: "TestSection", Key: "Key1", Value: "Value1"}
	jb, _ := json.Marshal(jj)
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/api/config", bytes.NewReader(jb))
	rr.ServeHTTP(w2, req2)

	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/api/config/Test/TestSection/Key1", nil)
	rr.ServeHTTP(w3, req3)

	fmt.Println(" @@@@@@@@@@@@@ body : ", w3.Body.String())
	fmt.Println("test unit test ")

	w4 := httptest.NewRecorder()
	req4, _ := http.NewRequest("DELETE", "/api/config/Test/TestSection/Key1", nil)
	rr.ServeHTTP(w4, req4)

	var cc models.ConfigProperty
	if err := json.Unmarshal(w4.Body.Bytes(), &cc); err != nil {
		fmt.Println("---------------------------------------- panic ")
		panic(err)
	}
	fmt.Println("Property : ", cc.ID)
	assert.Equal(t, cc.ID, uint(0), "프로퍼티를 삭제하지 못했습니다.")

}
