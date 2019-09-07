package diagram

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"downsampling"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	_ "gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CovertToPlotXY(data []downsampling.Point) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = data[i].X
		pts[i].Y = data[i].Y
	}
	return pts
}

func MakeLinePlotter(d plotter.XYs, c color.RGBA) (*plotter.Line, error) {
	// Make a line plotter and set its style.
	line, err := plotter.NewLine(d)
	if err != nil {
		return nil, err
	}
	line.LineStyle.Width = vg.Points(1)
	//rawLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	line.LineStyle.Color = c

	return line, nil
}

func SavePNG(title string, file string, name []string, line []*plotter.Line) error {
	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for i := range line {
		p.Add(line[i])
		p.Legend.Add(name[i], line[i])
	}
	if err := p.Save(18*vg.Inch, 6*vg.Inch, file); err != nil {
		return err
	}
	return nil
}

func ConcatPNGs(fileNames []string, targetFile string) error {

	images := make([]image.Image, len(fileNames))

	width := 0
	height := 0

	for i, f := range fileNames {
		file, err := os.Open(f)
		if err != nil {
			return err
		}
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
	for i := range images {
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
	return nil
}
