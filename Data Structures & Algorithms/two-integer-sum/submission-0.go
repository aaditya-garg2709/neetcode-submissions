func twoSum(nums []int, target int) []int {
   prevMap := make (map[int]int)
	// nums := [2,7,11,15]
	// target := 9
   for i,v := range nums{
	diff := target - v
	if j,found := prevMap[diff]; found{
		return []int{j,i}
	}
	prevMap[v] = i
   }
   return []int{}
}