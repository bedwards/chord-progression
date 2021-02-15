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
    "time"
)

var nextChord = map[int][]int{
	1: {2, 3, 4, 5, 6},
	2: {4, 5},
	3: {2, 4, 6},
	4: {1, 3, 5},
	5: {1, 4, 6},
	6: {1, 2, 4, 5},
}

type Numerals struct{}

func New() Numerals {
	rand.Seed(time.Now().UnixNano())
	return Numerals{}
}

type Part = []int

type SongABA struct {
	A Part
	B Part
}

const PART_LENGTH int = 16

const PICK_1_FIRST int = 60  // pick I first 60% of the time
const PICK_6_FIRST int = 30  // pick iv first 30% of the time

const REPEAT_FOUR_BARS int = 50 // if the four bars can be repeated, then repeat 50% of the time

const STAY_ON_1 int = 30  // stay on I 30% of the time
const STAY_ON_6 int = 15  // stay on iv 15% of the time
const STAY_ON_OTHER int = 3  // stay on other chords 3% of the time

func (n Numerals) pickFirst() int {
	// Prefer starting with 1 or 6
	val := rand.Intn(100)
	if val < PICK_1_FIRST {
		return 1
	}
	if val < PICK_1_FIRST + PICK_6_FIRST {
		return 6
	}
	others := []int{2, 3, 4, 5}
	return others[rand.Intn(len(others))]
}

func (n Numerals) pickNext(cur int) (int, error){
	choices, ok := nextChord[cur]
	if !ok {
		return -1, fmt.Errorf("Illegal chord numeral: %#v", cur)
	}
	if cur == 1 && rand.Intn(100) < STAY_ON_1 {
		return cur, nil
	} else if cur == 6 && rand.Intn(100) < STAY_ON_6 {
		return cur, nil
	} else if rand.Intn(100) < STAY_ON_OTHER {
		return cur, nil
	}
	return choices[rand.Intn(len(choices))], nil
}

func (n Numerals) createPart(first int) (Part, error) {
	// A part of a song. 16 bars long. E.g. verse, prechorus, chorus.
	var part []int
	part = append(part, first)
	for len(part) < PART_LENGTH {

		if len(part) > 0 && len(part) % 4 == 0 {
			last_four := part[len(part)-4:len(part)]
			if n.Verify(last_four, last_four) && rand.Intn(100) < REPEAT_FOUR_BARS {
				part = append(part, last_four...)
				continue
			}
		}

		next, err := n.pickNext(part[len(part)-1])
		if err != nil {
			return nil, err
		}
		part = append(part, next)

	}
	return part, nil
}

func (n Numerals) CreatePartThatFollows(last int) (Part, error) {
	next, err := n.pickNext(last)
	if err != nil {
		return nil, err
	}
	return n.createPart(next)
}

func (n Numerals) CreatePart() (Part, error) {
	return n.createPart(n.pickFirst())
}

func (n Numerals) Verify(left []int, right []int) bool {
	// Verify that right progression can follow left
	// Assumes that left and right are valid chord progressions
	last := left[len(left)-1]
	if right[0] == last {
		return true
	}
	for _, n := range nextChord[last] {
		if right[0] == n {
			return true
		}
	}
	return false
}

func (n Numerals) CreateSongABA() (*SongABA, error) {
	// A song with two parts. The parts can go back and forth legally
	// (A can follow B and B can follow A).
	a, err := n.CreatePart()
	if err != nil {
		return nil, fmt.Errorf("Failed to create song: %#v", err)
	}
	var b Part
	for keepTrying := true; keepTrying; keepTrying = !n.Verify(b, a) {
		b, err = n.CreatePartThatFollows(a[PART_LENGTH-1])
		if err != nil {
			return nil, fmt.Errorf("Failed to create song: %#v", err)
		}
	}
	return &SongABA{a, b}, nil
}
