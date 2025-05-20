package ritesh

import "fmt"

func Test1(x string) string {
	test2(x)
	return x
}

func test2(x string) {
	fmt.Println("From Pratik : ", x)
}
