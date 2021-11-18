package main

import "testing"

func TestEvent(t *testing.T) {
	t.Run("test max event should return 4", func(t *testing.T) {
		arrival := []int32{1, 3, 3, 5, 7}
		duration := []int32{2, 2, 1, 2, 1}

		result := maxEvent(arrival, duration)
		if result != 4 {
			t.Errorf("Expected 4, got %d", result)
		}
	})

	t.Run("test max event should return 1", func(t *testing.T) {
		arrival := []int32{1, 2}
		duration := []int32{7, 3}

		result := maxEvent(arrival, duration)
		if result != 1 {
			t.Errorf("Expected 1, got %d", result)
		}
	})

	t.Run("test max event should return 2", func(t *testing.T) {
		arrival := []int32{1, 3, 4, 6}
		duration := []int32{4, 3, 3, 2}

		result := maxEvent(arrival, duration)
		if result != 2 {
			t.Errorf("Expected 2, got %d", result)
		}
	})

}
