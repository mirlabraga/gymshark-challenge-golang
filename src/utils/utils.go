package utils

type Utils struct {
}

func (Utils) GetMaxElement(array []int, element int) (maxElement int) {
	maxElement = array[0]
	for i := 0; i < len(array); i++ {
		if element >= array[i] {
			maxElement = array[i]
		} else {
			break
		}
	}
	return maxElement
}

func (Utils) Contains(array []int, element int) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}
