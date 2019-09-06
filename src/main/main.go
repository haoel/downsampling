package main

import (
	"bufio"
	"encoding/csv"
	_ "fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"image"
	"image/color"
	"image/draw"
	"image/png"

	"downsampling"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	_ "gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func randFloatsPoints(min, max float64, n int) []downsampling.Point {
	rand.Seed(time.Now().UnixNano())
	res := make([]downsampling.Point, n)
	for i := range res {
		res[i].X = float64(i)
		res[i].Y = min + rand.Float64()*(max-min)
	}
	return res
}

func loadPointsFromCSV(file string) []downsampling.Point {
	csvFile, err := os.Open(file)
	checkError("Cannot Open the file.", err)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var data []downsampling.Point
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		checkError("Read file error", err)
		var d downsampling.Point
		d.X, _ = strconv.ParseFloat(line[0], 64)
		d.Y, _ = strconv.ParseFloat(line[1], 64)
		data = append(data, d)
	}
	return data
}

func covertToPlotXY(data []downsampling.Point) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = data[i].X
		pts[i].Y = data[i].Y
	}
	return pts
}

func savePNG(title string, file string, name []string, line []*plotter.Line) {
	p, err := plot.New()
	checkError("Create plot error", err)
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for i, _ := range line {
		p.Add(line[i])
		p.Legend.Add(name[i], line[i])
	}
	if err := p.Save(18*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}
}

func appendPNGs(fileNames []string, targetFile string) {

	images := make([]image.Image, len(fileNames))

	width := 0
	height := 0

	for i, f := range fileNames {
		file, err := os.Open(f)
		checkError("Cannot open file", err)
		images[i], _, err = image.Decode(file)
		if width < images[i].Bounds().Dx() {
			width = images[i].Bounds().Dx()
		}
		height += images[i].Bounds().Dy()
	}

	//rectangle for the big image
	rect := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}

	//create the new Image file
	rgba := image.NewRGBA(rect)

	height = 0
	for i, _ := range images {
		rect := images[i].Bounds().Add(image.Point{0, height})

		draw.Draw(rgba, rect, images[i], image.Point{0, 0}, draw.Src)
		height += images[i].Bounds().Dy()
	}

	// Encode as PNG.
	f, _ := os.Create(targetFile)
	png.Encode(f, rgba)
	f.Close()

	for _, f := range fileNames {
		os.Remove(f)
	}

}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkError("Get Binary directory error", err)
	dataDir := dir + "/../data/"

	rawdata := loadPointsFromCSV(dataDir + "source.csv")
	samples := downsampling.LTTB(rawdata, 500)

	// Make a line plotter and set its style.
	rawLine, err := plotter.NewLine(covertToPlotXY(rawdata))
	checkError("new line error", err)
	rawLine.LineStyle.Width = vg.Points(1)
	//rawLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	rawLine.LineStyle.Color = color.RGBA{R: 255, A: 255}

	sampleLine, err := plotter.NewLine(covertToPlotXY(samples))
	checkError("new line error", err)
	sampleLine.LineStyle.Width = vg.Points(1)
	//sampleLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	sampleLine.LineStyle.Color = color.RGBA{B: 255, A: 255}

	savePNG("Raw Data", "01.png", []string{"Raw Data"}, []*plotter.Line{rawLine})
	savePNG("DownSampling Data", "02.png", []string{"DownSampling Data"}, []*plotter.Line{sampleLine})
	savePNG("DownSampling Example", "03.png", []string{"Raw", "Sampled"}, []*plotter.Line{rawLine, sampleLine})

	appendPNGs([]string{"01.png", "02.png", "03.png"}, dataDir+"downsampling.chart.png")
}
