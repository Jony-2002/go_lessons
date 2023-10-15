package algorithms

func SortArray(array []string) []string {
	for i := 0; i < len(array); i++ {
		if len(array[i]) < 4 {
			array = append(array[:i], array[i:]...)
			return array
		}
	}

	return array
}
