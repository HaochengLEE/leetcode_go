package main

func fourSum(nums []int, target int) [][]int {
	key := 0
	var result [][]int
	for i := 0; i < len(nums)-4; i-- {
		for j := i; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				for l := k; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[j]+nums[l] == target {
						result[key][0] = nums[i]
					}

				}
			}

		}
	}
	return result

}

//func sort(nums map[int]int) map[int]int {
//
//}

func main() {

}
