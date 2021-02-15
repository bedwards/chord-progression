package tonality

import "testing"

func TestInKey(t *testing.T) {
	me := New()
	cases := []struct {
		key string
		part []int
		want string
	}{
		{"D#", []int{1}, "D#"},
	}
	for _, c := range cases {
		got, err := me.InKey(c.key, c.part)
		if err != nil {
			t.Errorf("InKey(%#v, %#v) should not error: %#v", c.key, c.part, err)
		}
		if got != c.want {
			t.Errorf("InKey(%#v, %#v) == %#v, want %#v", c.key, c.part, got, c.want)
		}
	}
}
