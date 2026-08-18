package bufgrow

import "testing"

// TestExtendSpareCapacityRegression is the regression test for the inkkeep07
// append bug: when dst had spare capacity, append(dst, extra) reused dst's
// backing array, so writes to the returned slice polluted the source slice.
// Extend must always return an independent slice.
func TestExtendSpareCapacityRegression(t *testing.T) {
	src := make([]byte, 3, 32) // len 3, cap 32 -> plenty of spare capacity
	copy(src, []byte{1, 2, 3})

	got := Extend(src, 9)

	// Shape and content of the result.
	if len(got) != 4 {
		t.Fatalf("len(got) = %d, want 4", len(got))
	}
	want := []byte{1, 2, 3, 9}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %d, want %d", i, got[i], want[i])
		}
	}

	// Writes to any index of the result must not write through to src,
	// neither into the visible portion nor into the spare capacity.
	got[0] = 77
	got[3] = 88
	if src[0] != 1 || src[1] != 2 || src[2] != 3 {
		t.Fatalf("Extend aliased src backing array: %v", src)
	}

	// Independence holds across multiple writes on the result.
	got[1] = 55
	if src[1] != 2 {
		t.Fatalf("second write aliased src: %v", src)
	}
}
