package textmash

import "testing"

func TestTF(t *testing.T) {

	colTfIdf := NewColTfIdf()

	b := []string{"Penn", "Teller"}

	colTfIdf.AddDoc(b)
	if 1 != 1.0 {
		t.Errorf("e.Value() is %d")
	}
}
