package diagram

import (
	"downsampling"
	"image/color"
)

type Config struct {
	Title string
	Name  string
	Data  []downsampling.Point
	Color color.RGBA
}
