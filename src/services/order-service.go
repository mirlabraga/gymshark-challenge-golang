package services

import (
	"math"
	"sort"

	"github.com/mirlabraga/gymshark-challenge-golang/src/models"
	"github.com/mirlabraga/gymshark-challenge-golang/src/utils"
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

	var maxElement = utils.Utils{}.GetMaxElement(packages, pack)
	var quantity = (int)(pack / maxElement)
	items = AddElement(items, quantity, maxElement)

	var waste int = pack % maxElement

	for waste > 0 {
		if pack <= packages[0] {
			items = AddElement(items, 1, packages[0])
		} else {
			var maxElementWaste int = utils.Utils{}.GetMaxElement(packages, waste)
			var newWaste = waste % maxElementWaste
			if newWaste == 0 {
				quantity = int(math.Ceil(float64(waste) / float64(maxElementWaste)))
			} else {
				quantity = 1
			}

			items = AddElement(items, quantity, maxElementWaste)

			if waste < maxElementWaste {
				break
			}
			waste = int(waste % maxElementWaste)
		}
	}

	return items
}

func ReducePackages(items []models.Item) (result []models.Item) {

	if len(items) == 1 {
		return items
	}

	sortElementAsc(items)

	var i int = 0
	for i < len(items)-1 {
		var sum int = int(items[i].Package + items[i+1].Package)
		var containsSum bool = utils.Utils{}.Contains(packages, sum)
		if containsSum {
			items = RemoveIndex(items, 0)
			items = RemoveIndex(items, 0)
			items = AddElement(items, 1, sum)
			i = 0
		} else {
			return items
		}
	}

	return items
}

func sortElementAsc(items []models.Item) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Package < items[j].Package
	})
}

func AddElement(items []models.Item, quantity int, packageValue int) (result []models.Item) {
	var item models.Item = models.Item{Quantity: quantity, Package: packageValue}
	items = append(items, item)
	return items
}

func RemoveIndex(array []models.Item, index int) []models.Item {
	ret := make([]models.Item, 0)
	ret = append(ret, array[:index]...)
	return append(ret, array[index+1:]...)
}
