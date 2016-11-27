package LongMath

func Add(a []int, b []int) []int {
	maxlen := Max(len(a), len(b))
	a = ChangeLength(a, maxlen)
	b = ChangeLength(b, maxlen)
	sum := make([]int, maxlen)
	carry := 0
	for i := 0; i < maxlen; i++ {
		sum[i], carry = Adder(a[i], b[i], carry)
	}
	if carry != 0 {
		sum = append(sum, carry)
	}
	return sum
}

func Adder(a, b, prev int) (int, int) {
	sum := a + b + prev
	return sum % 10, sum / 10

}

func AddStr(a string, b string) string{
	aneg, a := IsNegative(a)
	bneg, b := IsNegative(b)
	aint := ConverToArray(a)
	bint := ConverToArray(b)
	greater := Compare(aint, bint)
	if aneg == bneg {
		rez := Add(aint, bint)
		if(aneg){
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
	if(aneg && greater)||(bneg && !greater){
		return "-"+ConverToString(rez)
	}
	return ConverToString(rez)
}

func AddStrParallel(a string, b string, n int) string{
	aint := ConverToArray(a)
	bint := ConverToArray(b)
	rez := AddParallel(aint, bint, n)
	return ConverToString(rez)
}


func AddParallel(a []int, b []int, n int) []int{
	maxlen := Max(len(a), len(b))
	partlen := maxlen / n + 1
	maxlen = partlen * n
	a = ChangeLength(a, maxlen)
	b = ChangeLength(b, maxlen)
	sum := make([]int, maxlen)

	chanels := make([]chan  int, n+1)
	chanels[0] = make(chan int)
	for i:=0; i<n; i++{
		chanels[i+1] = make(chan int)
		go ParallelAdder(a[i*partlen:(i+1)*partlen], b[i*partlen:(i+1)*partlen], sum[i*partlen:(i+1)*partlen], chanels[i], chanels[i+1])
	}
	chanels[0] <- 0
	<-chanels[n]
	for sum[maxlen-1] == 0 {
		maxlen--
	}
	return sum[:maxlen]
}

func ParallelAdder(a, b, rez []int, prevChan, nextChan chan int){
	maxlen := Max(len(a), len(b))
	carry := 0
	for i := 0; i < maxlen; i++ {
		rez[i], carry = Adder(a[i], b[i], carry)
	}
	prevCarry := <-prevChan
	i := 0
	for prevCarry != 0 && i < maxlen{
		rez[i], prevCarry = Adder(rez[i], prevCarry, 0)
		i++
	}
	nextChan <- prevCarry + carry

}
