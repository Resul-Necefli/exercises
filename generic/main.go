package main

import "fmt"

func sumInt(a, b int) int {
	return a + b
}

func sumFLoat64(a, b float64) float64 {
	return a + b
}

/* burada bu yazi xetddi biraz menasiz gorunsede esas nuans odurki any  yeni isdedyin tipde ferq etmir  deyr gondere bilersen amma
bele tipler istfade edilen funksiya daxilnde emliyatlar adeten mehdudlasdirilir  cunkis  biz ora toplama yazsaq any her hansisa  bir deyer
ola bilceyinen proqram bize sehv vercey ki  men  bilmirem hansi tipde deyer gelcey ona gorede toplamani yerine yetire bilmerem   any adeten
fmt.Println(a ...any)  kimi gore  bilerik  yeni isdedyin qeder isdedyin tipde deyer gir ekrana yzadiracaq  adaeten bunukitapxnalarin ve ya hansisa
metodun yazilmasindadaha cox rast gele bilerik
*/

func genericType[T any](a, b T) (T, T) {
	return a, b
}

//	bu sadece gostermek ucun yazilan bir funckisyadir ve biraz menasizdir  https://go.dev/blog/comparable  bu url den biz comprable haqda
//
// daha genis melumat elde ede bileriy amma   bunu deyimki bu nov bir  nov muqayise  edile bilen tipleri muqayise ede bilmek ucun
// istfade edilir
func genericComprable[T comparable](a T) T {
	return a
}
func main() {
	fmt.Println("float funksiyasinan gelen deyer :", sumFLoat64(44.4, 55.2))
	fmt.Println("int funksiyasinan gelen deyr", sumInt(45, 77))
	a, b := genericType(445, 665)
	fmt.Println("genric funksiyaydan gelen deyerler:", a, b)
	c, d := genericType(445.444, 665.55)
	fmt.Println("genric funksiyaydan gelen deyerler:", c, d)

}
