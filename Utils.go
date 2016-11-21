package LongMath

import ("strconv")

func ChangeLength(input []int, len int) []int {
	output := make([]int, len)
	copy(output, input)
	return output
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Compare(a, b []int) bool {
	if len(a) > len(b) {
		return true
	}
	if len(a) < len(b) {
		return false
	}
	for i := len(a) - 1; i >= 0; i++ {
		if a[i] > b[i] {
			return true
		}
		if a[i] < b[i] {
			return false
		}
	}
	return true
}

func ConverToArray(str string) []int {
	var arr []int
	var element int

	for i := 0; i < len(str); i++ {
		element,_ = strconv.Atoi(string([]rune(str)[len(str)-i-1]))
		arr = append(arr, element)
	}

	return arr
}

func ConverToString(arr []int) string {
	var str string
	var i int = len(arr)-1

	for i >= 0  {
		str += strconv.Itoa(arr[i])
		i--
	}

	return str
}

func IsNegative(str string)(bool, string){
	if str[0] == '-'{
		return true, str[1:]
	}
	return false, str
}
