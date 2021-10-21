package builder

import "testing"

func TestNormalBuilder(t *testing.T) {
	// Create a new builder
	builder := getBuilder("normal")
	director := newDirector(builder)
	// Build the product
	house := director.construct()
	// Check the product
	if house.show() != "House: Wooden Window, Wooden Door, 1 floor\n" {
		t.Error("Expected wooden window")
	}
}
