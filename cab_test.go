package cab

import "testing"

func TestNewCabParser(t *testing.T) {
	_, err := NewCabParser("./cab-go-test.cab")
	if err != nil {
		t.Errorf("%s\n", err)
	}
}
