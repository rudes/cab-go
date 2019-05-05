package cab

import "testing"

func TestNewCabinet(t *testing.T) {
	_, err := NewCabinet("./cab-go-test.cab")
	if err != nil {
		t.Errorf("%s\n", err)
	}
}
