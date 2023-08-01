package SortUtils

import (
	"sort"
	"strconv"
)


func SortIntegers (list *[]string){
	sort.SliceStable(*list, func(i, j int) bool {
		var a, _ = strconv.Atoi((*list)[i])
		var b, _ = strconv.Atoi((*list)[j])
		return a < b
	})
}

func compareStringsASCII(str1, str2 string) bool {

	minLen := len(str1)

	if len(str2) < minLen {
		minLen = len(str2)
	}

	for i := 0; i < minLen; i++ {
		if int(str1[i]) < int(str2[i]) {
			return true
		}
	}
	return false

}

func SortMixs (list *[]string){
	sort.SliceStable(*list, func(i, j int) bool {
		return compareStringsASCII((*list)[i], (*list)[j])
	})
}

func SortStrings (list *[]string){
	sort.Strings(*list)
}

type SortFunc func(*[]string)