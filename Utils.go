package LongMath

func ChangeLength(input []int, len int) []int{
	output := make([]int, len)
	copy(output, input)
	return output
}

func Max(a int, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}

func Compare(a, b []int) bool{
	if len(a)>len(b){
		return true
	}
	if len(a)<len(b){
		return false
	}
	for i:= len(a)-1; i>=0; i++ {
		if(a[i]>b[i]){
			return true
		}
		if(a[i]<b[i]){
			return false
		}
	}
	return true
}

