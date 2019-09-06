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

	sampled := make([]Point, 0, threshold)

	// Bucket size. Leave room for start and end data points
	bucketSize := float64(len(data)-2) / float64(threshold-2)

	sampled = append(sampled, data[0]) // Always add the first point

	// We have 3 pointers represent for
	// > bucketLow - the current bucket's start location
	// > bucketMiddle - the current bucket's end location,
	//                  also the start location of next bucket
	// > bucketHight - the next bucket's end location.
	bucketLow := 1
	bucketMiddle := int(math.Floor(bucketSize)) + 1

	var prevMaxAreaPoint int

	for i := 0; i < threshold-2; i++ {

		bucketHigh := int(math.Floor(float64(i+2)*bucketSize)) + 1

		// Calculate point average for next bucket (containing c)
		avgRangeStart := bucketMiddle
		avgRangeEnd := bucketHigh

		if avgRangeEnd >= len(data) {
			avgRangeEnd = len(data)
		}

		avgRangeLength := float64(avgRangeEnd - avgRangeStart)

		var avgX, avgY float64
		for ; avgRangeStart < avgRangeEnd; avgRangeStart++ {
			avgX += data[avgRangeStart].X
			avgY += data[avgRangeStart].Y
		}
		avgX /= avgRangeLength
		avgY /= avgRangeLength

		// Get the range for current bucket
		currBucketStart := bucketLow
		currBucketEnd := bucketMiddle

		// Point a
		pX := data[prevMaxAreaPoint].X
		pY := data[prevMaxAreaPoint].Y

		maxArea := -1.0

		var maxAreaPoint int
		for ; currBucketStart < currBucketEnd; currBucketStart++ {
			// Calculate triangle area over three buckets
			area := (pX-avgX)*(data[currBucketStart].Y-pY) -
				(pX-data[currBucketStart].X)*(avgY-pY)
			// We only care about the relative area here.
			// Calling math.Abs() is slower than squaring
			area *= area
			if area > maxArea {
				maxArea = area
				maxAreaPoint = currBucketStart // Next a is this b
			}
		}

		sampled = append(sampled, data[maxAreaPoint]) // Pick this point from the bucket
		prevMaxAreaPoint = maxAreaPoint               // This MaxArea point is the next's prevMAxAreaPoint

		bucketLow = bucketMiddle
		bucketMiddle = bucketHigh
	}

	sampled = append(sampled, data[len(data)-1]) // Always add last

	return sampled
}
