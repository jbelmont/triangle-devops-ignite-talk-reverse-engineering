package hammingDistance

import (
	"errors"
	"fmt"
	"log"
)

func hammingDistance(x string, y string) (int, error) {
	if len(x) != len(y) {
		return 0, errors.New("Strings must of the same length")
	}
	var distance int

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			distance += 1
		}
	}

	return distance, nil
}

func main() {
	dist, err := hammingDistance("karolin", "kathrin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dist)
}
