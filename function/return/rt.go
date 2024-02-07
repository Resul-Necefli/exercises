package main

import "fmt"

/*
function  ozleride  bir  nov olduqlari ucun burada  func() func()  return  etmeye baxaq
*/
func rio() func() int {
	return func() int {
		return 45
	}
}

func main() {
	//  tek  ferq geri qayadan deyerin func olmasdir  buna  gorede biz
	y := rio()
	//  burada  y() moterzeler elave ederey onu ise  saliriq  yeni  rio dan bir  func deyer dondu
	//  y deyiskeninde  saxladim  ve  sonra  y  func  xasellerine sahip oldugu  ucun  ve  o  func geri donecek bir deyeri oldgu ucun
	// onu  ekreana  yazdirdim      biz   x:=   y()  de  yaza bilerik cunki  bu  func  da  geri deyer  donur
	y()
	fmt.Println(y())

	/*
		indi biz  func icnde function  parmetr olaraq verek  yeni artiq  bir  funksiyamiz  varki o  parmetr olaraq funksiya qebul edir
		amma  oz  sertlerine uygun gelen  funksiyani  yeni  interface  kimi  eyer  bu  funksiya  2 int  deyer  alir  ve  geri  1  int  deyer  qaytarirsa
		demeli   arqument  olaraq  gozdediyi  funksiyada  bu  teleblere uygun  olmadlir  eks teqdirde  biz  sehv  aliriq

	*/
	fmt.Println("------------------------------------")
	s := add(22, 44, collection)
	d := add(64, 55, subtract)

	ff := subtract(56, 55)
	gg := collection(55, 66)
	fmt.Println(ff)
	fmt.Println(gg)
	fmt.Println("------------------------------------")
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", collection)
	fmt.Printf("%T\n", subtract)
	fmt.Println("-------------------------------------")
	fmt.Println(s)
	fmt.Println(d)

}

//      func  icnde  func  parmetr  kimi verlirsede   untmyaq ki bu yalniz  anonim  func  ola  biler  cunki  func  icnde  func yalniz
//anonim olarq  qebul  edilir   ve  biz  burada   f(a,b)
/*
burada f(a,b)  nece  gorur?  nece  qebul edir ?   nece  yerindece oz-ozune  ise  dusur? axi func zeng edilmelidir
burda  f     add() funksiyasinin  icndedir gorduymuz  kimi    f  add()  funksiyasini  ozune    bir  nov  main()  funksiyasi kimi  gorur
cunki onun  kecerli oldgu tek  yer   bu  funksiyanin  icidir  ve  o  add()  funksiyasinan  ozune  a ve  b - ni rahtliqla  qebul edir
neceki biz  main() icinden basqa  bir  func a  arqumetler  oture biliriy  buda  ele     ve  bu  f()  function  eyni adresde
 ___saxlanilimir__   cunki o refrans ederey  add()  funca    baglanir  ve  onun  icndeki  a  ve  b  ni ozune  arqumet kimi  qebul edir
 ve  add() onun  ucun bir  nov main() func  evez  etdiyi  ucun  o  return f(a,b)   yazlmagi yetrlidirki  ise  dussun cu nki  o  oz  ana  funksiyasinin
 icnde zennedir ve arqumentlerini  qebul ederey ise  dusur
*/
func add(a int, b int, f func(int, int) int) int {
	return f(a, b)

}

func collection(a int, b int) int {

	return a

}

func subtract(a int, b int) int {
	return a - b
}
