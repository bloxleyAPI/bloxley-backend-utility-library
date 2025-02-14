package utilslibs

// Helper function to dereference a *string safely
func SafeString(s *string) string {
	if s == nil {
		return "" // or a default value
	}
	return *s
}

// String to pointer
func StringToPointer(s string) *string {
	return &s
}

// Bool to pointer
func BoolToPointer(b bool) *bool {
	return &b
}

// Contains returns true if slice s contains the element e.
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
