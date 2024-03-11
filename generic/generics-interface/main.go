package main

import (
	"fmt"
)

/*
genericler daha cix tekrarin qarsini almaq ucun istfade olunur  hansiki biz ferqli tipde olan arqumentlerimizle eyni emliyyari yerine
yetirirk  burda  2 func yazmagimiz mutleqdir amma buna gelin genricler nece baxir diqqet edek
*/
func sumFLoat64(a, b float64) float64 {
	return a + b
}

func sumInt(a, b int) int {
	return a + b
}

// generic tipimizi  teyin edrik sintax uygun olaraqdan ve ona hansi type leri ozunde saxlayacaginideyirik artqi biz yartdigimiz bu yeni novle
// birlikde  funksiyamiza 1 tipde deyil  bir nece tipde arqument oture bilery
func genericType[T int | float64](a, b T) T {

	return a + b

}

// biz hemcinin bu  isdedyimiz tipleri interface de teyin ederek artiq gneric tipi ve interfeyis destinde istfade etmis oluruq bu sadece
//
//	sade interface  mentiqi ile isleyir
type result interface {
	~float64 | ~int
}

func generInterface[T result](a, b T) T {
	return a + b
}

type n int

func main() {
	// biz artiq deyrlerimizi oture bilerik
	var a n = 4666
	fmt.Println("interface func gelen deyer ", generInterface(a, 55))
	fmt.Println(sumFLoat64(45.2, 48.6))
	fmt.Println(sumInt(78, 56))
	fmt.Println(genericType(77.77, 77.98))
	fmt.Println(genericType(44, 55))
}
