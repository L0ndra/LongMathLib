package LongMath


func Minus(a, b []int) []int {
	maxlen := len(a)
	a = ChangeLength(a, maxlen)
	b = ChangeLength(b, maxlen)
	diff := make([]int, maxlen)
	carry := 0
	for i := 0; i < maxlen; i++ {
		diff[i], carry = Differ(a[i], b[i], carry)
	}
	for diff[maxlen-1] == 0 {
		maxlen--
	}
	return diff[:maxlen]
}

func Differ(a, b, carry int) (int, int) {
	return (a - b + carry + 10) % 10, (a - b + carry - 10) / 10
}

func MinusStr(a, b string) string {
	aneg, a := IsNegative(a)
	bneg, b := IsNegative(b)
	aint := ConverToArray(a)
	bint := ConverToArray(b)
	greater := Compare(aint, bint)
	if aneg != bneg {
		rez := Add(aint, bint)
		if aneg {
			return "-"+ConverToString(rez)
		}
		return ConverToString(rez)
	}
	var rez []int;
	if greater {
		rez = Minus(aint, bint)
	}else{
		rez = Minus(bint, aint)
	}
	if(aneg && greater)||(!bneg && !greater){
		return "-"+ConverToString(rez)
	}
	return ConverToString(rez)
}
