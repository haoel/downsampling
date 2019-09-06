package downsampling

import (
	"math"
)

func calculateTriangleArea (pa, pb, pc Point) float64 {
	area := ((pa.X - pc.X) * (pb.Y - pa.Y) - (pa.X - pb.X) * (pc.Y - pa.Y)) * 0.5
	return math.Abs(area)
}


func calculateAverageDataPoint(points []Point) (avg Point) {

	for _ , point := range points {
		avg.X += point.X
		avg.Y += point.Y
	}
	l := float64(len(points))
	avg.X /= l
	avg.Y /= l
	return avg
}