package _3_strconv

import (
	"fmt"
	"strconv"
)

func Test() {
	num, _ := strconv.Atoi("20")
	fmt.Println(num)

	num, _ = strconv.Atoi("2")
	fmt.Println(num)

	num64, _ := strconv.ParseInt("2000000000000000", 10, 64)
	fmt.Println(num64)

	str := strconv.Itoa(20)
	fmt.Printf("%q\n", str)

	str = strconv.FormatInt(2000000000000000, 10)
	fmt.Printf("%q\n", str)
}
