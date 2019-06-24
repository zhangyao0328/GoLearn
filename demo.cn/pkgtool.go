package main

import (
	"GoLearnProject/mode"
	"fmt"
)

var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
}

func init() {

	ins := []int{3, 3, 1, 5, 3, 5, 6, 8, 9}

	mp := make(map[int]string)

	mp[1] = "tom"
	mp[1] = "pony"
	mp[2] = "jaky"
	mp[3] = "ands"

	goto L1

	delete(mp, 5)

	for i, v := range ins {
		fmt.Println(i, v)
	}

	fmt.Println("***************")

L1:
	user := mode.User{Name:"yao"}
	fmt.Println(user.Name, user.Arg)
	//c := map[string]int{"a" : 1,"b" : 2}

}
