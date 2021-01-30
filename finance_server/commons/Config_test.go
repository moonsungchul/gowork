package commons_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/moonsungchul/finance/commons"
	"github.com/tkanos/gonfig"
)

func Test1(t *testing.T) {
	conf := commons.Config{}

	err := gonfig.GetConf("../conf/config.json", &conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
	assert.Equal(t, conf.MySQL_Dbname, "fms_finance")
}
