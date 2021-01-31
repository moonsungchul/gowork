package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/moonsungchul/finance/commons"
	"github.com/moonsungchul/finance/controllers"
	"github.com/moonsungchul/finance/models"
)

func setupRouter() *gin.Engine {
	config := commons.NewConfig("../conf/config.json")
	store := models.MysqlStore{}
	gold := controllers.RestGold{Config: config, Store: &store}
	rr := gin.Default()
	rr.GET("/api/gold", gold.GetGoldData)
	rr.GET("/api/gold/graph", gold.GetGoldDataGraph)
	return rr
}

func TestRestGold(t *testing.T) {
	rr := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/gold?start=0&rows=100", nil)

	rr.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)

}

func TestRestGoldGraph(t *testing.T) {
	rr := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/gold/graph", nil)

	rr.ServeHTTP(w, req)
	//fmt.Println(w.Body)
	assert.Equal(t, 200, w.Code)
}
