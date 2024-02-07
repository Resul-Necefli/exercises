package main

import "fmt"

func foo() {
	fmt.Println("foo")
}

// Bar funksiyasini yardaq
func Bar() {
	fmt.Println("bar")
}

func main() {
	/*
	   func main() icnde  proqram ise  dusdukde yuxaridan  asgi  dogru irelileyir bezi  halarda  biz   func  icnde
	   her  hansi  bir  fayil  aciriq   os.open   ve  ya  database  qosularken baglanti  yaradarken   biz
	   bezi  emliyatlari  baglamali  oluruq  acdigimiz  fayl  kimi    ve  bu  kimi  emliyatlarin  sonda  olmasini
	   yeni func  icndeki  en  sonuncu  emliyat  olaraq  yerine  yetirlmesini  isdeyirkse  o  zaman  biz  burada   DEFER   ifadesini istfade
	   etmeliyik

	*/
	defer foo()
	Bar()

	//  burada  vurlmag isdediyim  cixis   0 ,4,3,2,1  kimi  olcaq  proqram yuxardan asgi gelir  ve  defer ifadeleri ucun
	//  birde yuxariya  dogru  qaydir bu zaman  hersey bir birnin tersine  cep  olunur yeni defer() ifadeleri olanlardan sohbet  gedir
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("0")

}
