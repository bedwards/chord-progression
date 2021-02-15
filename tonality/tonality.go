package tonality

import (
	"fmt"
	"strings"
	"container/ring"
)

var allNotes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

type Tonality struct {
	chromatic *ring.Ring
}

func New() Tonality {
	tonality := Tonality{}
	tonality.chromatic = ring.New(len(allNotes))
	for i := 0; i < len(allNotes); i++ {
		tonality.chromatic.Value = allNotes[i]
		tonality.chromatic = tonality.chromatic.Next()
	}
	return tonality
}

func (t Tonality) InKey(key string, part []int) (string, error) {
	for _, _ = range allNotes {
		if t.chromatic.Value == key {
			break
		}
		t.chromatic = t.chromatic.Next()
	}
	if t.chromatic.Value != key {
		return "", fmt.Errorf("Illegal key: %#v", key)
	}
	chords := make(map[int]string)
	chords[1] = fmt.Sprintf("%-3v", t.chromatic.Value.(string))
	t.chromatic = t.chromatic.Next()
	t.chromatic = t.chromatic.Next()
	chords[2] = fmt.Sprintf("%-3v", t.chromatic.Value.(string) + "m")
	t.chromatic = t.chromatic.Next()
	t.chromatic = t.chromatic.Next()
	chords[3] = fmt.Sprintf("%-3v", t.chromatic.Value.(string) + "m")
	t.chromatic = t.chromatic.Next()
	chords[4] = fmt.Sprintf("%-3v", t.chromatic.Value.(string))
	t.chromatic = t.chromatic.Next()
	t.chromatic = t.chromatic.Next()
	chords[5] = fmt.Sprintf("%-3v", t.chromatic.Value.(string))
	t.chromatic = t.chromatic.Next()
	t.chromatic = t.chromatic.Next()
	chords[6] = fmt.Sprintf("%-3v", t.chromatic.Value.(string) + "m")
	var b strings.Builder
	for i, j := range part {
		if j < 1 || j > 6 {
			return "", fmt.Errorf("Illegal part: %#v", part)
		}
		if i > 0 && i % 4 == 0 {
			b.WriteString("\n")
		}
		b.WriteString(chords[j] + " ")
	}

	return strings.TrimRight(b.String(), " "), nil
}
