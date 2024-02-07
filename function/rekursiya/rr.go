package main

import "fmt"

//  burada  sadece olarak  rekursiya yeni ic ice olan birseydi biz  bunu  for  ilede  ede  biliriy
// bu eynen  faktorial  kimdir

func description(b int) int {

	if b == 0 {
		return 1
	} // faktorial mentiqi ile isleyir
	// b 4 deyeri  gelir  ve  bura  olur  bele  olur  4 *(4-1) = 4*3   sonraki  deyer  3 *(3-1)
	return b * description(b-1)
}

func main() {

	c := description(4)
	fmt.Println(c)
	fmt.Println(forloop(5))

}

// bunu  for  dongesi ile  edek  daha  asan ve  basa  duslur olur   ve  biz  burada  5 faktorialini  hesablamis  oluruq

func forloop(v int) int {

	i := 1
	for v > 0 {
		i *= v
		v--
	}

	return i

}
