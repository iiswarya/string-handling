// interface + struct + ==

package main

import "fmt"

type user struct {
	Name string
}

type userDetail struct {
	Name string
	Tags []string
}

func main() {
	var a interface{} = user{Name: "Nick"}
	var b interface{} = user{Name: "Nick"}

	fmt.Println(a == b) // true because go compares dynamic types and values (type - user, value - Nick)

	var c interface{} = userDetail{Name: "Nick", Tags: []string{"admin", "user"}}
	var d interface{} = userDetail{Name: "Nick", Tags: []string{"admin", "user"}}

	fmt.Println(c == d) // false because slice, maps and functions are not comparable. so non-comparable value inside an interface{} doesn't make it comparable.

}
