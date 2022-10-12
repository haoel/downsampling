package diagram

import (
	"github.com/haoel/downsampling/core"
	"image/color"
)

type Config struct {
	Title string
	Name  string
	Data  []core.Point
	Color color.RGBA
}
