type Solution struct {
	Key string
}

func (s *Solution) Encode(strs []string) string {
	// If the slice is empty, return a specific marker that wouldn't happen otherwise
	if len(strs) == 0 {
		return "EMPTY_ARRAY_MARKER"
	}
	return strings.Join(strs, "🍏")
}

func (s *Solution) Decode(e string) []string {
	// Catch our specific empty marker and return a completely empty slice
	if e == "EMPTY_ARRAY_MARKER" {
		return []string{}
	}
	if len(e) == 0 {
		return []string{""} // Handles the case where the input was [""]
	}
	return strings.Split(e, "🍏")
}