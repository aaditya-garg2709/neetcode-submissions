func topKFrequent(nums []int, k int) []int {
    count := make(map [int]int)
	for _,v := range nums{
		count[v]++
	}
	freq := make([][]int, len(nums)+1)	
	for i,v := range count{
		freq[v] = append(freq[v],i)
	}
	topKFreq := []int{}
	for i := len(freq) - 1; i > 0; i--{
		for _,v := range freq[i]{
			topKFreq = append(topKFreq,v)
			if len(topKFreq) == k{
				return topKFreq
			}
		}
	}
	return topKFreq
}