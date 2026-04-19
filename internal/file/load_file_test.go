package file

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	expected := []byte{0xFF, 0xA0, 0x00, 0x1B}

	tmp, err := os.CreateTemp("", "rom-*.ch8")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(expected); err != nil {
		t.Fatal(err)
	}
	tmp.Close()

	got, err := ReadFile(tmp.Name())
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	if len(got) != len(expected) {
		t.Fatalf("length: want %d, got %d", len(expected), len(got))
	}
	for i, b := range expected {
		if got[i] != b {
			t.Errorf("byte[%d]: want 0x%02X, got 0x%02X", i, b, got[i])
		}
	}
}
