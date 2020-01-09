package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	biggestListMp := make(map[int]bool, len(nums1))
	var smallestList []int
	var biggestList []int

	if len(nums1) < len(nums2) {
		smallestList = nums1
		biggestList = nums2
	} else {
		smallestList = nums2
		biggestList = nums1
	}

	result := make([]int, 0, len(smallestList))
  intersectionMp := make(map[int]bool, len(smallestList))

	for _, val := range biggestList {
		biggestListMp[val] = true
	}

	for _, val := range smallestList {
		if _, saw := biggestListMp[val]; saw {
      intersectionMp[val] = true
		}
	}

  for ky := range intersectionMp {
    result = append(result, ky)
  }

	return result
}

func main() {
	fmt.Printf("intercepted %v\n", intersection([]int{61, 24, 20, 58, 95, 53, 17, 32, 45, 85, 70, 20, 83, 62, 35, 89, 5, 95, 12, 86, 58, 77, 30, 64, 46, 13, 5, 92, 67, 40, 20, 38, 31, 18, 89, 85, 7, 30, 67, 34, 62, 35, 47, 98, 3, 41, 53, 26, 66, 40, 54, 44, 57, 46, 70, 60, 4, 63, 82, 42, 65, 59, 17, 98, 29, 72, 1, 96, 82, 66, 98, 6, 92, 31, 43, 81, 88, 60, 10, 55, 66, 82, 0, 79, 11, 81}, []int{5, 25, 4, 39, 57, 49, 93, 79, 7, 8, 49, 89, 2, 7, 73, 88, 45, 15, 34, 92, 84, 38, 85, 34, 16, 6, 99, 0, 2, 36, 68, 52, 73, 50, 77, 44, 61, 48}))
}
