package main

import (
	"bufio"
	"encoding/csv"
	"image/color"
	"io"
	"os"
	"strconv"

	"common"
	"diagram"
	"downsampling"
)

func loadPointsFromCSV(file string) []downsampling.Point {
	csvFile, err := os.Open(file)
	common.CheckError("Cannot Open the file.", err)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var data []downsampling.Point
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		common.CheckError("Read file error", err)
		var d downsampling.Point
		d.X, _ = strconv.ParseFloat(line[0], 64)
		d.Y, _ = strconv.ParseFloat(line[1], 64)
		data = append(data, d)
	}
	return data
}

func main() {

	dir := common.GetBinaryDirectory()
	dataDir := dir + "/../data/"

	const sampledCount = 500
	rawdata := loadPointsFromCSV(dataDir + "source.csv")
	samplesLTOB := downsampling.LTOB(rawdata, sampledCount)
	samplesLTTB := downsampling.LTTB(rawdata, sampledCount)
	samplesLTD := downsampling.LTD(rawdata, sampledCount)

	var dcs = []diagram.Config{
		{Title: "Raw Data", Name: "Raw Data", Data: rawdata, Color: color.RGBA{A: 255}},
		{Title: "LTOB Sampled Data", Name: "Sampled - LTOB", Data: samplesLTOB, Color: color.RGBA{R: 255, A: 255}},
		{Title: "LTTB Sampled Data", Name: "Sampled - LTTD", Data: samplesLTTB, Color: color.RGBA{B: 255, A: 255}},
		{Title: "LTD Sampled Data", Name: "Sampled - LTD", Data: samplesLTD, Color: color.RGBA{G: 255, A: 255}},
	}

	err := diagram.CreateDiagram(dcs, dataDir+"downsampling.chart.png")
	common.CheckError("create diagram error", err)

}
