// Chord progressions are usually expressed by Roman numerals in classical
// music theory. In rock and blues, musicians also often refer to chord
// progressions using Roman numerals, as this facilitates transposing a song to
// a new key.

// This package deals with expression chord progressions using numerals
// (1=I, 2=ii, 3=iii, 4=IV, 5=V, 6=vi). The 7th diminished chord isn't used.

// This package does not deal with the key (tonality) of the song. Therefore
// chord progressions in this package are not expressed using the name and
// "quality" (e.g. Dm) of the chords.

package numerals

import (
	"fmt"
	"math/rand"
)

var nextChord = map[int][]int{
	1: {2, 3, 4, 5, 6},
	2: {4, 5},
	3: {2, 4, 6},
	4: {1, 3, 5},
	5: {1, 4, 6},
	6: {1, 2, 4, 5},
}

const PART_LENGTH int = 16

func pickFirst() int {
	// Prefer starting with 1 or 6
	val := rand.Intn(100)
	if val < 60 {
		return 1
	}
	if val < 90 {
		return 6
	}
	others := []int{2, 3, 4, 5}
	return others[rand.Intn(len(others))]
}

func pickNext(cur int) (int, error){
	choices, ok := nextChord[cur]
	if !ok {
		return -1, fmt.Errorf("Illegal chord numeral: %#v", cur)
	}
	return choices[rand.Intn(len(choices))], nil
}

func CreatePart() ([]int, error) {
	var part []int
	part = append(part, pickFirst())
	for len(part) < PART_LENGTH {
		next, err := pickNext(part[len(part)-1])
		if err != nil {
			return nil, err
		}
		part = append(part, next)
	}
	return part, nil
}
