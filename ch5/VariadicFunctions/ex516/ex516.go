package ex516

import "fmt"

func StringJoin(ss ...string) (join string) {
	shouldReturn, returnValue := newFunction(ss)
	if shouldReturn {
		return returnValue
	}

	for _, s := range ss {
		join += s
	}

	return
}

func newFunction(ss []string) (bool, string) {
	if len(ss) == 0 {
		fmt.Println("require at least one arguement")
		return true, ""
	}
	return false, ""
}
