package diagram

import (
	"downsampling/core"
	"image/color"
)

type Config struct {
	Title string
	Name  string
	Data  []core.Point
	Color color.RGBA
}
