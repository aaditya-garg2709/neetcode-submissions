func hasDuplicate(nums []int) bool {
    indi := make(map [int]int)
    for _,v := range nums{
        indi[v]+=1
        if indi[v] >1{
            return true
        }
    }
    return false
}
