package LongMath

func Minus(a, b []int) []int{
	maxlen := Max(len(a), len(b))
	isPossitive := Compare(a, b)
	if !isPossitive{
		a, b = b, a
	}
	a = ChangeLength(a, maxlen)
	b = ChangeLength(b, maxlen)
	diff := make([]int, maxlen)
	carry := 0
	for i:=0; i<maxlen; i++{
		diff[i], carry = Differ(a[i], b[i], carry)
	}
	for diff[maxlen] == 0 {
		maxlen--
	}
	return diff
}

func Differ(a, b, carry int) (int, int){
	return 	(a - b + carry + 10) % 10, (a - b + carry - 10)/10
}
