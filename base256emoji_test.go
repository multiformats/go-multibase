package multibase

import "testing"

func TestBase256EmojiAlphabet(t *testing.T) {
	var c uint
	for _, v := range base256emojiTable {
		if v != rune(0) {
			c++
		}
	}
	if c != 256 {
		t.Errorf("Base256Emoji count is wrong, expected 256, got %d.", c)
	}
}

func TestBase256EmojiUniq(t *testing.T) {
	m := make(map[rune]struct{}, len(base256emojiTable))
	for i, v := range base256emojiTable {
		_, ok := m[v]
		if ok {
			t.Errorf("Base256Emoji duplicate %s at index %d.", string(v), i)
		}
		m[v] = struct{}{}
	}
}
