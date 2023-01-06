package sorts

type Sorter interface {
	Len() int
	Swap(i, j int)
}

func SelectionSort(items Sorter, less func(i, j int) bool) {
	n := items.Len()
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if less(j, minIndex) {
				minIndex = j
			}
		}
		items.Swap(i, minIndex)
	}
}
