package core_test

import (
	"os"
	"testing"

	"github.com/haoel/downsampling/core"
	"github.com/haoel/downsampling/demo/common"
)

func BenchmarkLTTB(b *testing.B) {
	dir, _ := os.Getwd()
	dataDir := dir + "/../demo/data/"

	const sampledCount = 500
	rawdata := common.LoadPointsFromCSV(dataDir + "source.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		core.LTTB(rawdata, sampledCount)
	}
}

func BenchmarkLTTB2(b *testing.B) {
	dir, _ := os.Getwd()
	dataDir := dir + "/../demo/data/"

	const sampledCount = 500
	rawdata := common.LoadPointsFromCSV(dataDir + "source.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		core.LTTB2(rawdata, sampledCount)
	}
}
