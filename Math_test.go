package LongMath

import "testing"
import "math/big"
import "math/rand"
import "strconv"

func TestAdd1(t *testing.T){
	rez := AddStr("333","999")
	if(rez != "1332"){
		t.Error("Rezult %s not equal 1332",rez)
	}
}

func TestAdd2(t *testing.T){
	rez := AddStr("-333","999")
	if(rez != "666"){
		t.Error("Rezult %s not equal 666",rez)
	}
}

func TestAdd3(t *testing.T){
	rez := AddStr("333","-999")
	if(rez != "-666"){
		t.Error("Rezult %s not equal -666",rez)
	}
}

func TestAdd4(t *testing.T){
	rez := AddStr("-333","-999")
	if(rez != "-1332"){
		t.Error("Rezult %s not equal -666",rez)
	}
}

func TestMinus1(t *testing.T)  {
	rez := MinusStr("333", "999")
	if(rez != "-666"){
		t.Error("Rezult %s not equal -666",rez)
	}
}

func TestMinus2(t *testing.T)  {
	rez := MinusStr("-333", "999")
	if(rez != "-1332"){
		t.Error("Rezult %s not equal -1332",rez)
	}
}

func TestMinus3(t *testing.T)  {
	rez := MinusStr("333", "-999")
	if(rez != "1332"){
		t.Error("Rezult %s not equal 1332",rez)
	}
}

func TestMinus4(t *testing.T)  {
	rez := MinusStr("-333", "-999")
	if(rez != "666"){
		t.Error("Rezult %s not equal 666",rez)
	}
}

func TestAddLong(t *testing.T){
	var a string = ""
	len := rand.Int()%200000
	for i := 0; i < len; i++{
		a = a + strconv.Itoa(rand.Int()%10)
	}
	var b string = ""
	len = rand.Int()%200000
	for i := 0; i < len; i++{
		b = b + strconv.Itoa(rand.Int()%10)
	}

	rez := AddStr(a,b)
	var testA, testB, testRez big.Int
	testA.SetString(a,10)
	testB.SetString(b,10)
	testRez.Add(&testA, &testB)

	if rez != testRez.String(){
		t.Error("Rezult %s not equal %s", rez, testRez.String())
	}
}





