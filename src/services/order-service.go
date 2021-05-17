package services

import (
	"math"
	"sort"

	"github.com/mirlabraga/gymshark-challenge-golang/src/models"
)

var packages []int = []int{250, 500}

func Calculation(pack int) (items []models.Item) {

	if pack <= packages[0] {
		var item = models.Item{Quantity: 1, Package: packages[0]}
		items = append(items, item)
		return items
	}

	items = Decomposition(pack)
	items = reducePackages(items)
	return
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
			//return items

			if waste < maxValueBeforeWaste {
				break
			}
			waste = waste % maxValueBeforeWaste
		}
	}

	return
}

func reducePackages(items []models.Item) (result []models.Item) {

	sort.Slice(items, func(i, j int) bool {
		return items[i].Package < items[j].Package
	})

	var i int = 0
	for i < len(items)-1 {
		var sum = items[i].Package + items[i+1].Package
		if contains(sum) {
			items = append(items[:i], items[i+1:]...)
			items = append(items[:i], items[i+1:]...)
			var item = models.Item{Quantity: 1, Package: sum}
			result = []models.Item{item}
			i = 0
		} else {
			break
		}
	}

	return
}

func RemoveIndex(s []models.Item, index int) []models.Item {
	print(s)
	return append(s[:index], s[index+1:]...)
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
	var i int = 0
	maxValueMinPack = packages[0]

	for pack >= packages[i] {
		maxValueMinPack = packages[i]
		i++
	}
	return
}
