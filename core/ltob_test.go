package core_test

import (
	"os"
	"testing"

	"github.com/haoel/downsampling/core"
	"github.com/haoel/downsampling/demo/common"
)

func BenchmarkLTOB(b *testing.B) {
	dir, _ := os.Getwd()
	dataDir := dir + "/../demo/data/"

	const sampledCount = 500
	rawdata := common.LoadPointsFromCSV(dataDir + "source.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		core.LTOB(rawdata, sampledCount)
	}
}
