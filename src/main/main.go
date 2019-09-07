package main

import (
	_ "fmt"
	"io"
	"os"
	"bufio"
	"strconv"
	"encoding/csv"
	"image/color"

	"common"
	"diagram"
	"downsampling"

	"gonum.org/v1/plot/plotter"
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

	rawdata := loadPointsFromCSV(dataDir + "source.csv")
	smaplesLTOB := downsampling.LTOB(rawdata, 500)
	smaplesLTTB := downsampling.LTTB(rawdata, 500)

	rawLine, err := diagram.MakeLinePlotter(diagram.CovertToPlotXY(rawdata), color.RGBA{R: 0, G: 0, B: 0, A: 255})
	common.CheckError("Cannot make a line plotter", err)

	sampleLineLTOB, err := diagram.MakeLinePlotter(diagram.CovertToPlotXY(smaplesLTOB), color.RGBA{R: 255, G: 0, B: 0, A: 255})
	common.CheckError("Cannot make a line plotter", err)

	sampleLineLTTB, err := diagram.MakeLinePlotter(diagram.CovertToPlotXY(smaplesLTTB), color.RGBA{R: 0, G: 0, B: 255, A: 255})
	common.CheckError("Cannot make a line plotter", err)

	if err := diagram.SavePNG("Raw Data", "01.png",
		[]string{"Raw Data"}, []*plotter.Line{rawLine}); err != nil {
		common.LogFatal("Cannot save the png file", err)
	}
	if err := diagram.SavePNG("DownSampling Data - LTOB", "02.png",
		[]string{"DownSampling Data - LTOB"}, []*plotter.Line{sampleLineLTOB}); err != nil {
		common.LogFatal("Cannot save the png file", err)
	}
	if err := diagram.SavePNG("DownSampling Data - LTTB", "03.png",
		[]string{"DownSampling Data - LTTB"}, []*plotter.Line{sampleLineLTTB}); err != nil {
		common.LogFatal("Cannot save the png file", err)
	}
	if err := diagram.SavePNG("DownSampling Example", "05.png",
		[]string{"Raw", "Sampled - LTOB", "Sampled - LTTB"},
		[]*plotter.Line{rawLine, sampleLineLTOB, sampleLineLTTB}); err != nil {
		common.LogFatal("Cannot save the png file", err)
	}

	if err := diagram.ConcatPNGs([]string{"01.png", "02.png", "03.png", "05.png"}, dataDir+"downsampling.chart.png"); err != nil {
		common.LogFatal("Cannot concatenate the png files", err)
	}
}
