package main

func maxEvent(arrival, duration []int32) int32 {
	var counter int = 1

	for i := 1; i < len(arrival); i++ {
		tempDuration := arrival[i-1] + duration[i-1] - 1
		if tempDuration <= arrival[i] {
			counter++
		}
	}

	return int32(counter)
}
