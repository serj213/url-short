package randomalias

import "testing"

func TestRandomAlias(t *testing.T) {
	length := 10
	alias := RandomAlias(length)

	if len(alias) != length {
		t.Errorf("Exepting length alias %d, result %d", length, len(alias))
	}

	for _, char := range alias {
		if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z') {
			t.Errorf("Alias contains invalid character: %c", char)
		}
	}
}

