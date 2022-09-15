package mypackage

import "sort"

func ValeursCentrales(sl []int) []int{
	sort.Ints(sl)
	r := make([]int,1)
	if len(sl)%2==1{
		r[0] = sl[len(sl)/2]
	}else {
		r[0] = sl[len(sl)/2-1]
		r = append(r,sl[len(sl)/2])
	}
	return r
}
