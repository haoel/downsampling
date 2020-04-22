package downsampling_test

import (
	"os"
	"testing"

	"common"
	"downsampling"
)

func BenchmarkLTD(b *testing.B) {
	dir, _ := os.Getwd()
	dataDir := dir + "/../../data/"

	const sampledCount = 500
	rawdata := common.LoadPointsFromCSV(dataDir + "source.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		downsampling.LTOB(rawdata, sampledCount)
	}
}
