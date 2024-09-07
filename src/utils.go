package breakout

func Every[T any](arr []T, cond func(elem T) bool) bool {
	for i := 0; i < len(arr); i += 1 {
		if !cond(arr[i]) {
			return false
		}
	}
	return true
}
