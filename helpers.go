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
