package utils

type StringSlice []string

// Called for clearing an array
func (s *StringSlice) Clear() {
	*s = make([]string, 0)
}

// Called for checking if an element is already apprended in the slice
func (s *StringSlice) ExistAlready(element string) bool {
	for _, e := range *s {
		if e == element {
			return true
		}
	}
	return false
}
