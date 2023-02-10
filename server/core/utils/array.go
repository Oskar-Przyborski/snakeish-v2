package utils

func ArrayIncludes[T comparable](arr []T, searchItem T) bool {
	for _, item := range arr {
		if item == searchItem {
			return true
		}
	}
	return false
}
