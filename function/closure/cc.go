/*
Bağlama daxili funksiya strukturu ilə əlaqəli bir anlayışdır.
Funksiya daxilində təyin olunan və xarici dəyişənlərə daxil olan daxili funksiyalar bu xarici dəyişənləri istinad kimi qəbul edir
və onları saxlayır. Bu daxili funksiyalar xarici funksiyanın işləmə müddətində yaradıldığından,
hətta xarici funksiyanın əhatə dairəsi bitdikdən sonra da onların xarici funksiyanın dəyişənlərinə çıxışı var.
 Bu mexanizm “bağlama” adlanır,
 çünki funksiyanın əhatə dairəsini daşıyan bu daxili funksiyalar xaricdən xarici funksiyanın dəyişənlərinə daxil ola bilir.

Bağlamaların istifadəsi funksional proqramlaşdırma üslubunda olduqca yaygındır və müxtəlif vəziyyətlərdə faydalıdır.
 Xüsusilə geri çağırış funksiyaları, funksional parametrlər və dəyişən əhatə dairəsi tələb olunan vəziyyətlərdə qapanmalar tez-tez istifadə olunur.


*/

package main

import "fmt"

func ttx(a, b int) func() int {
	c := a + b
	return func() int {
		c++
		return c
	}

}

func context() func() int {

	x := 0
	return func() int {
		x++
		return x
	}

}

func main() {

	resizster := ttx(44, 55)
	for i := 0; i < 5; i++ {

		fmt.Println(resizster())
	}

	f := context()
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("------------------------------------------")
	// burada  gordymuz kimi  herbir  deyer ferqli deyiskenlerde saxlandigi  kimi  ferqli  adreslerde  tutulut
	// bunun ucun g  ile  zengi funksiyaya etdikde o yeniden baslayir
	g := context()
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())

}
