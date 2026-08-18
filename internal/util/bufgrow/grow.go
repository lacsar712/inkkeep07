package bufgrow

// Extend returns a new slice with extra appended. The result must not share
// a backing array with dst even when dst has spare capacity.
func Extend(dst []byte, extra byte) []byte {
	return append(dst, extra)
}

func Concat(a, b []byte) []byte {
	out := make([]byte, len(a)+len(b))
	copy(out, a)
	copy(out[len(a):], b)
	return out
}
