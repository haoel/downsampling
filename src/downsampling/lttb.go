package downsampling

import (
	"math"
)

// LTTB down-samples the data to contain only threshold number of points that
// have the same visual shape as the original data
func LTTB(data []Point, threshold int) []Point {

	if threshold >= len(data) || threshold == 0 {
		return data // Nothing to do
	}

	sampledData := make([]Point, 0, threshold)

	// Bucket size. Leave room for start and end data points
	bucketSize := float64(len(data)-2) / float64(threshold-2)

	sampledData = append(sampledData, data[0]) // Always add the first point

	// We have 3 pointers represent for
	// > bucketLow - the current bucket's beginning location
	// > bucketMiddle - the current bucket's ending location,
	//                  also the beginning location of next bucket
	// > bucketHight - the next bucket's ending location.
	bucketLow := 1
	bucketMiddle := int(math.Floor(bucketSize)) + 1

	var prevMaxAreaPoint int

	for i := 0; i < threshold-2; i++ {

		bucketHigh := int(math.Floor(float64(i+2)*bucketSize)) + 1

		// Calculate point average for next bucket (containing c)
		avgPoint := calculateAverageDataPoint(data[bucketMiddle : bucketHigh+1])

		// Get the range for current bucket
		currBucketStart := bucketLow
		currBucketEnd := bucketMiddle

		// Point a
		pointA := data[prevMaxAreaPoint]

		maxArea := -1.0

		var maxAreaPoint int
		for ; currBucketStart < currBucketEnd; currBucketStart++ {

			area := calculateTriangleArea(pointA, avgPoint, data[currBucketStart])
			if area > maxArea {
				maxArea = area
				maxAreaPoint = currBucketStart
			}
		}

		sampledData = append(sampledData, data[maxAreaPoint]) // Pick this point from the bucket
		prevMaxAreaPoint = maxAreaPoint                       // This MaxArea point is the next's prevMAxAreaPoint

		//move to the next window
		bucketLow = bucketMiddle
		bucketMiddle = bucketHigh
	}

	sampledData = append(sampledData, data[len(data)-1]) // Always add last

	return sampledData
}
