package util

func Contains[T comparable](source []T, findBy T) bool {
	for _, item := range source {
		if item == findBy {
			return true
		}
	}
	return false
}
