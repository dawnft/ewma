package ewma

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	avg1 := Average()
	if avg1 != 0.0 {
		t.Errorf("Must be 0.0 but %f", avg1)
	}

	avg2 := Average(1, 2, 3)
	if avg2 != 2.0 {
		t.Errorf("Must be 2.0 but %f", avg2)
	}

	avg3 := Average(-5, 5, -10.1, 10.1, 5.0)
	if avg3 != 1.0 {
		t.Errorf("Must be 1.0 but %f", avg3)
	}

	avg4 := Average(.34, 232.5, 23.2, 1.4, 565.1, 34.1, 123.9, -0.1)
	if fmt.Sprintf("%3.2f", avg4) != "122.55" {
		t.Errorf("Must be 122.55 but %.2f", avg4)
	}

	avg5 := Average(23, -2.1, 2.4, 4.1, 3.52, 6.1, 7.1)
	if fmt.Sprintf("%1.5f", avg5) != "6.30286" {
		t.Errorf("Must be 6.30286 but %.5f", avg5)
	}
}
