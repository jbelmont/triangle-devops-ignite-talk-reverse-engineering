package hammingDistance

import "testing"

func TestHammingDistance(t *testing.T) {
	dst, err := hammingDistance("x", "xx")
	if dst != 0 && err.Error() != "Strings must of the same length" {
		t.Error("distance should be zero and should receive an error when strings are of different length")
	}

	dist, err := hammingDistance("x", "x")
	if err != nil {
		t.Error("Should be nil")
	}
	if dist != 0 {
		t.Error("Hamming Distance should be 0")
	}

	dist, err = hammingDistance("x", "y")
	if err != nil {
		t.Error("Should be nil")
	}
	if dist != 1 {
		t.Error("Hamming Distance should be 1")
	}

	dist, err = hammingDistance("karolin", "kerstin")
	if err != nil {
		t.Error("Should be nil")
	}
	if dist != 3 {
		t.Error("Hamming Distance should be 3")
	}
}
