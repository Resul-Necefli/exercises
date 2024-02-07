package main

import "fmt"

//  ilkin  olaraq  standart  bir variadic  funk yazaq

func sum(i ...int) int {

	var total int

	for _, v := range i {

		total += v
	}
	return total

}

func main() {

	s := sum(1, 2, 3, 4, 5, 8, 10)
	fmt.Println(s)

	p := pluarValues(sum, 1, 2, 3, 4, 5, 8, 10, 15, 16, 17, 18)
	fmt.Println(p)
	a := singluarValues(sum, 2, 3, 4, 5, 8, 10, 15, 16, 17, 18)
	fmt.Println(a)

}

// burada  func  parametr alan func goruruk  daha  oncede  gorduk indi ise  tek  ferq   bu  icindeki anonim  func variadicdir ve
//
//	o  arqument olaraqda aldigi  funksiya ozunde bu kimi funksionaliqlari  olmaldi  ki bu  func onu  qebul ede  bilsin
func pluarValues(f func(xi ...int) int, d ...int) int {

	var total []int

	for _, v := range d {
		if v%2 == 0 {
			total = append(total, v)
		}

	}
	return f(total...)
}

func singluarValues(f func(yi ...int) int, c ...int) int {

	var sign []int

	for _, g := range c {

		if g%2 != 0 {

			sign = append(sign, g)
		}

	}
	//  burada   sum()  funksiyasi  f  argumentidir  eslinde  o  total deyisekenindeki deyerleri toplayir  bir  int olaraq  geri qaytarir
	//  f isdediyi  kimi  yeni
	return f(sign...)
}
