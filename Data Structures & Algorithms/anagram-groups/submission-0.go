func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0{
		return nil
	}
	var ana [][]string
	list := make (map [[26]int][]string)
	for _, v := range strs{
		alpha := [26]int{}
		for _,x := range v{
			alpha[x - 'a']++
		}
		list[alpha] = append(list[alpha],v)
	}
	for _,v := range list{
		ana = append(ana,v)
	}
	return ana
}