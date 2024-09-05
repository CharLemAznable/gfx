package gx

func SliceMapping[T any, R any](from []T, mapping func(t T) R) (to []R) {
	for _, t := range from {
		to = append(to, mapping(t))
	}
	return
}

func SliceInterfaces[T any](from []T) []interface{} {
	return SliceMapping(from, func(t T) interface{} { return t })
}
