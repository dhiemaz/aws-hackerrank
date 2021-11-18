package main

import "testing"

func TestPromo(t *testing.T) {
	t.Run("testing foo -- must 1", func(t *testing.T) {
		var codeList []string = []string{"apple apple", "banana anything banana"}
		var shoppingCart []string = []string{"orange", "apple", "apple", "banana", "orange", "banana"}
		result := foo(codeList, shoppingCart)

		if result != 1 {
			t.Error("Expected 1, got ", result)
		}
	})

	t.Run("testing foo -- must 0", func(t *testing.T) {
		var codeList []string = []string{"apple apple", "apple apple banana"}
		var shoppingCart []string = []string{"apple", "apple", "apple", "banana"}
		result := foo(codeList, shoppingCart)

		if result != 0 {
			t.Error("Expected 0, got ", result)
		}
	})

	t.Run("testing foo -- must 0", func(t *testing.T) {
		var codeList []string = []string{"apple apple", "banana anything banana"}
		var shoppingCart []string = []string{"apple", "banana", "apple", "banana", "orange", "banana"}
		result := foo(codeList, shoppingCart)

		if result != 0 {
			t.Error("Expected 0, got ", result)
		}
	})
}
