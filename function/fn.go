//   func

package main

import (
	"fmt"
)

// 1  ilkin  olaraq  hecbir  parmater  almayan  bir  function yardaq    funksiyamizi   main()  funksiyasi  daxilinde  cagirdigda
//
//	o bize  sadece olaraq    yazdirmaq  isdediymiz  formati   gostercey
func foo() {
	fmt.Println("hello gopher")
}

// indi  ise     r adinda  ve  string type    parmetiri   elave  edek  ama bize  hecne  qaytarmasin return etmesin yeni
//
//	biz  artiq  ana  funksiyamiz  olan main function  dan  bura  arqument oturmeliyik  ve  o  arqumenti  bir  deyiskende saxlmaliyiq
//
// sebebin  func main  daxilinde  bu  func -a  zeng  etdikde  yazacam
func name(r string) {
	fmt.Println("menim adim", r)
}

//	burada  artix  funksiyamiza  verilen deyerlerle   birlikde  geri  bir  deyer  qaytaraq ve
//
// onun  hansi  tipde  olacagini  funca  bildirek
func result(s string) string {

	return "name:" + s
}

func text(f string, i int) (string, int) {

	return f, i
}

//firs_name, last_name, age := strings.Split(name, " ")

func main() {
	//   goda func  bu  sekilde cagirlir   ()  moterzeler   foo   funksiya  oldugunu  bildirir
	foo()
	//  gelin  bir  parametrler  elave  etdiymiz  funksiani  isledey
	// asgda  gordyumuz  kimi  yazma  sebebimiz  bizim  func hecbir  deyer  dondurmeycek
	//  sadce    buradan  o  func ise  salincaq   ve  o  function oz isini  bitrecek ve  functiondan   cixcaq ve proqram yoluna  davam  edecek
	//  cunki  biz  geriye  deyer  qaytarmir   eger  biz  bu  sekilde  yazsaq   a := name("resul")  func  hecbir  deyer qaytarmadigi  ucun
	//  go bize  isledmediymiz  bir  deyisken tanimladigimiz  ucun  erorr verir   bu  func sadce  oz icndeki  formati  yazdirir burada
	//  fmt.Println(result("orxan")) bu  sekilde  yazdirmis  olsaydiq  go  xeta  oldugunu  bildrecek  cunki eyer  funksiyadan
	//  bir  deyer  qayitmasa  bunu  yazdirmaq  menasiz  olardi
	name("Reusul")

	// indi  ise  biz  geri  deyer  qaytaran  bir  funksiyaya  bir  deyer  oturey
	//  artiq  biz  bu  prosesde  ordan gelen  deyeri bir  deyiskende   saxlamaliyiq

	a := result("orxan")
	fmt.Println(a) //  bu  formatda  yzdirma  sebebimiz  odurki  biz    gelen deyeri deyiskende  tutmaliyiq  bunun  tebiki
	//  ne ucun  etmiymizi  bir  sonraki  functionda  aciqlayaq

	e, b := text("James bond", 55)

	fmt.Println(e, b)

}
