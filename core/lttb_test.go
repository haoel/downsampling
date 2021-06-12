package core_test

import (
	"os"
	"testing"

	"downsampling/common"
	"downsampling/core"
)

func BenchmarkLTTB(b *testing.B) {
	dir, _ := os.Getwd()
	dataDir := dir + "/../data/"

	const sampledCount = 500
	rawdata := common.LoadPointsFromCSV(dataDir + "source.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		core.LTTB(rawdata, sampledCount)
	}
}
