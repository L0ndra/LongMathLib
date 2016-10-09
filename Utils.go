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
