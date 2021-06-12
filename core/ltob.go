package core

import (
	"math"
)

// Largest triangle one bucket(LTOB) data downsampling algorithm implementation
//  - Require: data . The original data
//  - Require: threshold . Number of data points to be returned
func LTOB(data []Point, threshold int) []Point {

	if threshold >= len(data) || threshold == 0 {
		return data // Nothing to do
	}

	sampledData := make([]Point, 0, threshold)

	// Bucket size. Leave room for start and end data points
	bucketSize := float64(len(data)-2) / float64(threshold-2)

	sampledData = append(sampledData, data[0]) // Always add the first point

	for bucket := 1; bucket < threshold-1; bucket++ {
		startIdx := int(math.Floor(float64(bucket) * bucketSize))
		endIdx := int(math.Min(float64(len(data)-1), float64(bucket+1)*bucketSize))

		maxArea := -1.0
		maxAreaIdx := -1

		for i := startIdx; i < endIdx; i++ {
			area := calculateTriangleArea(data[i-1], data[i], data[i+1])
			if area > maxArea {
				maxArea = area
				maxAreaIdx = i
			}

		}

		sampledData = append(sampledData, data[maxAreaIdx])
	}

	sampledData = append(sampledData, data[len(data)-1]) // Always add last

	return sampledData
}
