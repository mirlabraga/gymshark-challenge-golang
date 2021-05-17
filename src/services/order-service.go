package services

import (
	"math"
	"sort"

	"github.com/mirlabraga/gymshark-challenge-golang/src/models"
)

var packages []int = []int{250, 500, 1000, 2000, 5000}

func Calculation(pack int) (items []models.Item) {

	if pack <= packages[0] {
		var item = models.Item{Quantity: 1, Package: packages[0]}
		items = append(items, item)
		return items
	}

	items = Decomposition(pack)
	items = ReducePackages(items)
	return items
}

func Decomposition(pack int) (items []models.Item) {

	var maxValueBefore = getMaxValueBefore(pack)
	var quantity = (int)(pack / maxValueBefore)

	var item = models.Item{Quantity: quantity, Package: maxValueBefore}
	items = append(items, item)

	var waste int = pack % maxValueBefore

	for waste > 0 {
		if pack <= packages[0] {
			var item = models.Item{Quantity: 1, Package: packages[0]}
			items = append(items, item)
			return items
		} else {
			var maxValueBeforeWaste = getMaxValueBefore(waste)
			var newWaste = waste % maxValueBeforeWaste
			if newWaste == 0 {
				quantity = int(math.Ceil(float64(waste) / float64(maxValueBeforeWaste)))
			} else {
				quantity = 1
			}

			var item = models.Item{Quantity: quantity, Package: maxValueBeforeWaste}
			items = append(items, item)

			if waste < maxValueBeforeWaste {
				break
			}
			waste = int(waste % maxValueBeforeWaste)
		}
	}

	return items
}

func ReducePackages(items []models.Item) (result []models.Item) {

	if len(items) == 1 {
		return items
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Package < items[j].Package
	})

	var i int = 0
	for i < len(items)-1 {
		var sum int = int(items[i].Package + items[i+1].Package)
		if contains(sum) {
			items = append(items[:i], items[i+1:]...)
			items = append(items[:i], items[i+1:]...)
			var item = models.Item{Quantity: 1, Package: sum}
			items = append(items, item)
			i = 0
		} else {
			return items
		}
	}

	return items
}

func contains(element int) bool {
	for _, v := range packages {
		if v == element {
			return true
		}
	}
	return false
}

func getMaxValueBefore(pack int) (maxValueMinPack int) {

	maxValueMinPack = packages[0]
	for i := 0; i < len(packages); i++ {

		if pack >= packages[i] {
			maxValueMinPack = packages[i]
		} else {
			break
		}
	}
	return
}
