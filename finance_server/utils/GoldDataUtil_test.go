package utils_test

import (
	"testing"

	"github.com/moonsungchul/finance/utils"
)

func Test1(t *testing.T) {
	util := utils.NewGoldDataUtil("C:/work/data_workspace/finance/src/gold.csv")
	util.LoadCSV()
	/*
		util.CalCloseLastClose()
		util.CalVolumeLastVolume()
		util.CalCloseDivLastClose(5, "Close5LastClose")
		util.CalCloseDivLastClose(10, "Close10LastClose")
		util.CalCloseDivLastClose(20, "Close20LastClose")
		util.CalCloseDivLastClose(60, "Close60LastClose")
		util.CalCloseDivLastClose(120, "Close120LastClose")
		util.CalVolumeDivLastVolume(5, "Volume5LastVolume")
		util.CalVolumeDivLastVolume(10, "Volume10LastVolume")
		util.CalVolumeDivLastVolume(20, "Volume20LastVolume")
		util.CalVolumeDivLastVolume(60, "Volume60LastVolume")
		util.CalVolumeDivLastVolume(120, "Volume120LastVolume")
	*/
	util.ToCSV("gold_final.csv")
}
