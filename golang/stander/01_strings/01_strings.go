package _1_strings

import (
	"fmt"
	"strings"
)

func Test() {
	fmt.Println("\n包含子字符串:")
	fmt.Println(strings.Contains("abc123", "123"))
	fmt.Println(strings.ContainsRune("abc123", 'c'))

	fmt.Println("\n\n子串出现次数:")
	fmt.Println(strings.Count("fivevev", "vev"))
	fmt.Println(strings.Count("fivevev", "ve"))

	fmt.Println("\n分割字符串:")
	fmt.Printf("%q\n", strings.Fields("  foo bar  baz   "))
	fmt.Printf("%q\n", strings.FieldsFunc("aa,bb,cc", func(c rune) bool {
		if c == ',' {
			return true
		} else {
			return false
		}
	}))
	fmt.Printf("%q\n", strings.Split("aa,bb,cc", ","))
	fmt.Printf("%q\n", strings.SplitAfter("aa,bb,cc", ","))
	fmt.Printf("%q\n", strings.SplitN("aa,bb,cc", ",", 2))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println("\n判断字符串前后缀:")
	fmt.Println(strings.HasPrefix("sabc", "sa"))
	fmt.Println(strings.HasSuffix("sabc", "bc"))

	fmt.Println("\n字符串出现位置:")
	fmt.Println(strings.Index("sabcdefv", "bc"))
	fmt.Println(strings.IndexRune("sabcdefv", 'c'))
	fmt.Println(strings.LastIndex("sabcdefbcv", "bc"))

	fmt.Println("\n字符串拼接:")
	fmt.Println(strings.Join([]string{"aa", "bb", "cc", "dd"}, ", "))

	fmt.Println("\n字符串重复几次")
	fmt.Println(strings.Repeat("abc", 3))

	fmt.Println("\n字符串替换")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

	s := 1
	s, a := 2, 3
	fmt.Println(s, a)
}
