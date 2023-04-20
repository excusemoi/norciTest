package hexConverter

import (
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		hexString string
		decString string
	}{
		{
			hexString: "fff",
			decString: "4095",
		},
		{
			hexString: "abcd",
			decString: "43981",
		},
	}

	for _, c := range cases {
		result := Convert(c.hexString)
		if result != c.decString {
			t.Errorf("result %s != expected %s", result, c.decString)
		}
	}
}
