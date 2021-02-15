package numerals

import "testing"

type IntSet map[int]struct{}
var y = struct{}{}  // used to say 'yes' it is in the set

func TestPickNext(t *testing.T) {
	n := New()
	cases := []struct {
		in int
		want IntSet
	}{
		{1, IntSet{1:y, 2:y, 3:y, 4:y, 5:y, 6:y}},
		{2, IntSet{2:y, 4:y, 5:y}},
		{3, IntSet{2:y, 3:y, 4:y, 6:y}},
		{4, IntSet{1:y, 3:y, 4:y, 5:y}},
		{5, IntSet{1:y, 4:y, 5:y, 6:y}},
		{6, IntSet{1:y, 2:y, 4:y, 5:y, 6:y}},
	}
	for _, c := range cases {
		got, err := n.pickNext(c.in)
		if err != nil {
			t.Errorf("pickNext(%#v) == %#v, want %#v", c.in, err, c.want)
		}
		if _, ok := c.want[got]; !ok {
			t.Errorf("pickNext(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
	for _, i := range []int{0, 7} {
		if _, err := n.pickNext(i); err == nil {
			t.Errorf("pickNext(%#v) should error", i)
		}
	}
}

func TestCreatePart(t *testing.T) {
	n := New()
	part, err := n.CreatePart()
	if err != nil {
		t.Errorf("err should be nil: %#v", err)
	}
	if len(part) != 16 {
		t.Errorf("part should have 16 bars: %#v", len(part))
	}
}

func TestVerify(t *testing.T) {
	n := New()
	left := []int{1,1,1,1, 1,1,1,1, 1,1,1,1, 1,1,1,2}
	cases := []struct {
		right []int
		want bool
	}{
		{[]int{4,1,1,1, 1,1,1,1, 1,1,1,1, 1,1,1,1}, true},
		{[]int{3,1,1,1, 1,1,1,1, 1,1,1,1, 1,1,1,1}, false},
	}
	for _, c := range cases {
		got := n.Verify(left, c.right)
		if got != c.want {
			t.Errorf("Verify(%#v, %#v) == %#v, want %#v", left, c.right, got, c.want)
		}
	}
}

func TestCreateSongABA(t *testing.T) {
	n := New()
	song, err := n.CreateSongABA()
	if err != nil {
		t.Errorf("CreateSongABA() should not error: %#v", err)
	}
	if !n.Verify(song.A, song.B) {
		t.Errorf("CreateSongABA() b=%#v cannot follow a=%#v", song.B, song.A)
	}
	if !n.Verify(song.B, song.A) {
		t.Errorf("CreateSongABA() a=%#v cannot follow b=%#v", song.A, song.B)
	}
}
