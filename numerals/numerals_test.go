package numerals

import "testing"

type IntSet map[int]struct{}
var y = struct{}{}  // used to say 'yes' it is in the set

func TestPickNext(t *testing.T) {
	cases := []struct {
		in int
		want IntSet
	}{
		{1, IntSet{2:y, 3:y, 4:y, 5:y, 6:y}},
		{2, IntSet{4:y, 5:y}},
		{3, IntSet{2:y, 4:y, 6:y}},
		{4, IntSet{1:y, 3:y, 5:y}},
		{5, IntSet{1:y, 4:y, 6:y}},
		{6, IntSet{1:y, 2:y, 4:y, 5:y}},
	}
	for _, c := range cases {
		got, err := pickNext(c.in)
		if err != nil {
			t.Errorf("pickNext(%#v) == %#v, want %#v", c.in, err, c.want)
		}
		if _, ok := c.want[got]; !ok {
			t.Errorf("pickNext(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
	for _, i := range []int{0, 7} {
		if _, err := pickNext(i); err == nil {
			t.Errorf("pickNext(%#v) should error", i)
		}
	}
}

func TestCreatePart(t *testing.T) {
	part, err := CreatePart()
	if err != nil {
		t.Errorf("err should be nil: %#v", err)
	}
	if len(part) != 16 {
		t.Errorf("part should have 16 bars: %#v", len(part))
	}
}
