package main

import "testing"

func TestParking(t *testing.T) {
	t.Run("test parking roof should return 6", func(t *testing.T) {
		cars := []int{6, 2, 12, 7}
		k := 3

		min := carParkingRoof(cars, k)
		if min != 6 {
			t.Errorf("Expected 6, got %d", min)
		}
	})
}
