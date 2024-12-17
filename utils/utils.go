package utils

func Buscar[T any](slice []T, criterio func(item T) bool) bool {
	for _, item := range slice {
		if criterio(item) {
			return true
		}
	}
	return false
}
