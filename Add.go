package LongMath

func Add(a []int, b []int) []int {
	maxlen := Max(len(a), len(b))
	a = ChangeLength(a, maxlen)
	b = ChangeLength(b, maxlen)
	sum := make([]int, maxlen)
	carry := 0
	for i:=0; i<maxlen; i++{
		sum[i], carry = Adder(a[i], b[i], carry)
	}
	if(carry != 0) {
		sum = append(sum, carry)
	}
	return sum
}

func Adder(a, b, prev int) (int, int){
	sum := a + b + prev
	return sum % 10, sum / 10

}


