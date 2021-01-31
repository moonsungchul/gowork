package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moonsungchul/finance/commons"
	"github.com/moonsungchul/finance/models"
)

type RestGold struct {
	Config *commons.Config
	Store  *models.MysqlStore
}

func (r RestGold) GetGoldData(c *gin.Context) {
	start := c.Query("start")
	rows := c.Query("rows")
	print("conf : ", r.Config.MySQL_Dbname)
	print("start, rows", start, rows)
	db := r.Store.Open(r.Config.MySQL_Host, r.Config.MySQL_Port, r.Config.MySQL_Dbname,
		r.Config.MySQL_User, r.Config.MySQL_Passwd)
	total := r.Store.GetPagesTotal(db)
	istart, _ := strconv.Atoi(start)
	irows, _ := strconv.Atoi(rows)
	ar := r.Store.GetPricesPages(db, istart, irows)

	c.JSON(200, gin.H{
		"total": total,
		"start": start,
		"rows":  ar,
		"msg":   "return ok",
	})
}

func (r RestGold) GetGoldDataGraph(c *gin.Context) {
	db := r.Store.Open(r.Config.MySQL_Host, r.Config.MySQL_Port, r.Config.MySQL_Dbname,
		r.Config.MySQL_User, r.Config.MySQL_Passwd)
	ar := r.Store.GetPrices(db)
	c.JSON(200, gin.H{
		"rows": ar,
		"msg":  "return ok",
	})
}
