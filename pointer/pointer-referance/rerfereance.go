// pointerler de diger type ler  kimi  yeni  map slices interface ve s  kimi deyerini  yox  referansi paylasir
// buna dair bir nece meseleye baxaq

package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("xi deyeri -->", xi) //xi deyeri --> [1 2 3 4 5 6 7]
	xiUpdate(xi)
	fmt.Println("xi deyeri-->", xi) //xi deyeri--> [25 2 3 4 5 6 7]
	// burada  gordyumuz  kimi  slice   yarnanda onun altinda yarnan arrye  refrance  etdiyi  ucun  biz  bu slice basqa  deyiskenle tesir
	// edrik ve orginal  slice  de qiymeti  deyisir

	//  eynen xeritelerdede  beledir buna  dair misal yazaq
	a := make(map[string]int)
	a["James"] = 45
	fmt.Println("a deyeri", a)
	mapUpadte(a, "James")
	fmt.Println("a nin deyeri", a) //  a nin deyerinin  deysidiyini  gore bilrik cunki  xeritelerde  referansi paylasir  deyeri deyil
	/*
	 deyer ve  referance paylasma  mentiqi  biz  proqram  islyende  arxa terefde     bir varki meslen   i:=5    bu diyekki deyeri paylasir
	  proqram islyende i nin sadece  deyerinin kopyasin  alip yoluna  davam edecek   ama   referanse paylasilanda    i  nin adresini  ozunu
	  esas gotruru yeni birbasa  i  nin  yadasdaki yerini esas  gotur  buda ona getrip cixarirki  basqa  deyisende bu adres varsa
	  i  ye  tesir  ede  biler   pointrler  tam olaraq bu mentiqde  isleyir  yeni senin evinin mende  adresi varsa men ora gelde  bilerem
	  seni ordan cixarada bilerem    senin adresin basqa birnde varsa oda  bunlari ede biler
	*/
	c := 55
	fmt.Println("c nin deyeri --->", c)
	fmt.Println("c nin adresid -->", c)
	addPoint(&c)
	fmt.Println("c nin deyeri --->")

}

//  adres  , unvan qebul eden bir func  yazaq

func addPoint(v *int) { //  bu  funksiya   c adresin alip onun deyerine  5 elave  edip ve yeniden  c adresine geri gonderir belikle
	*v = *v + 5 //  c = 60 olur atix  cunki bu birbasa  yadasdaki yerden deyeri  alip  deysip yeniden ora qoyur
}

func xiUpdate(a []int) {
	a[0] = 25
}

func mapUpadte(m map[string]int, s string) {
	m[s] = 42
}
