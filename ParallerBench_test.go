package LongMath

import "testing"
import "math/rand"
import "strconv"

func BenchmarkParallel(t *testing.B){
	var a string = ""
	len := 100000
	for i := 0; i < len; i++{
		a = a + strconv.Itoa(rand.Int()%10)
	}
	var b string = ""
	len = 100000
	for i := 0; i < len; i++{
		b = b + strconv.Itoa(rand.Int()%10)
	}

	AddStrParallel(a,b,8)

}

func BenchmarkLinear(t *testing.B){
	var a string = ""
	len := 100000
	for i := 0; i < len; i++{
		a = a + strconv.Itoa(rand.Int()%10)
	}
	var b string = ""
	len = 100000
	for i := 0; i < len; i++{
		b = b + strconv.Itoa(rand.Int()%10)
	}

	AddStr(a,b)

}
