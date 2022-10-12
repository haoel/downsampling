package main

import (
	"flag"
	"image/color"
	"log"
	"os"
	"runtime/pprof"

	"github.com/haoel/downsampling/core"
	"github.com/haoel/downsampling/demo/common"
	"github.com/haoel/downsampling/demo/diagram"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	dir := common.GetBinaryDirectory()
	dataDir := dir + "/../data/"

	const sampledCount = 500

	log.Println("Reading the testing data...")
	rawdata := common.LoadPointsFromCSV(dataDir + "source.csv")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		var x []core.Point
		for i := 0; i < 200; i++ {
			x = core.LTOB(rawdata, sampledCount)
			x = core.LTTB(rawdata, sampledCount)
			x = core.LTD(rawdata, sampledCount)
		}
		pprof.StopCPUProfile()
		println("%v\n", x)
	}
	log.Printf("Downsampling the data from %d to %d...\n", len(rawdata), sampledCount)
	samplesLTOB := core.LTOB(rawdata, sampledCount)
	common.SavePointsToCSV(dataDir+"downsampling.ltob.csv", samplesLTOB)
	log.Println("Downsampling data - LTOB algorithm done!")
	samplesLTTB := core.LTTB(rawdata, sampledCount)
	common.SavePointsToCSV(dataDir+"downsampling.lttb.csv", samplesLTTB)
	log.Println("Downsampling data - LTTB algorithm done!")
	samplesLTD := core.LTD(rawdata, sampledCount)
	common.SavePointsToCSV(dataDir+"downsampling.ltd.csv", samplesLTD)
	log.Println("Downsampling data - LTD algorithm done!")

	file := dataDir + "downsampling.chart.png"
	log.Printf("Creating the diagram file...")
	var dcs = []diagram.Config{
		{Title: "Raw Data", Name: "Raw Data", Data: rawdata, Color: color.RGBA{A: 255}},
		{Title: "LTOB Sampled Data", Name: "Sampled - LTOB", Data: samplesLTOB, Color: color.RGBA{R: 255, A: 255}},
		{Title: "LTTB Sampled Data", Name: "Sampled - LTTD", Data: samplesLTTB, Color: color.RGBA{B: 255, A: 255}},
		{Title: "LTD Sampled Data", Name: "Sampled - LTD", Data: samplesLTD, Color: color.RGBA{G: 255, A: 255}},
	}

	err := diagram.CreateDiagram(dcs, file)
	common.CheckError("create diagram error", err)
	log.Printf("Successfully created the diagram - %s\n", file)
}
