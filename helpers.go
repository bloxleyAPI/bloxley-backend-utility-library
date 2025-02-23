package utilslibs

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

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

// String to Dob 2006-01-02
func StringToDobMust(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}

// GetMustPersonID returns the person_id from the context.
func GetMustPersonID(c context.Context, key string) uuid.UUID {
	personID := c.Value(key).(string)

	if personID == "" {
		panic(errors.New("person_id not found in context"))
	}
	personIDUUID := uuid.MustParse(personID)
	return personIDUUID

}
