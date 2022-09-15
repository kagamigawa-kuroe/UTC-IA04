package mypackage

func Moyenne(sl []int) int{
	sum := 0
	for i := range sl{
		sum = sum + i;
	}
	return sum/len(sl);
}
