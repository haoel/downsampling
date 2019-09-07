package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"image/color"
	"io"
	"log"
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

func savePointsToCSV(file string, points []downsampling.Point) {
	fp, err := os.Create(file)
	common.CheckError("Cannot create file", err)
	defer fp.Close()

	writer := csv.NewWriter(fp)
	defer writer.Flush()

	for _, point := range points {
		x := fmt.Sprintf("%f", point.X)
		y := fmt.Sprintf("%f", point.Y)
		err := writer.Write([]string{x, y})
		common.CheckError("Cannot write to file", err)
	}
}

func main() {

	dir := common.GetBinaryDirectory()
	dataDir := dir + "/../data/"

	const sampledCount = 500

	log.Println("Reading the testing data...")
	rawdata := loadPointsFromCSV(dataDir + "source.csv")

	log.Printf("Downsampling the data from %d to %d...\n", len(rawdata), sampledCount)
	samplesLTOB := downsampling.LTOB(rawdata, sampledCount)
	savePointsToCSV(dataDir+"downsampling.ltob.csv", samplesLTOB)
	log.Println("Downsampling data - LTOB algorithm done!")
	samplesLTTB := downsampling.LTTB(rawdata, sampledCount)
	savePointsToCSV(dataDir+"downsampling.lttb.csv", samplesLTTB)
	log.Println("Downsampling data - LTTB algorithm done!")
	samplesLTD := downsampling.LTD(rawdata, sampledCount)
	savePointsToCSV(dataDir+"downsampling.ltd.csv", samplesLTD)
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
