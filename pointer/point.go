/*
pointer demeli bizim   deyerlerimizin  yadasda bir deyeri birde  onun adresi olur   biz  bu  iki  seyle ordaki deyerin ne oldugunu gore bilirik

	meslen  yadasda bir yer var  orda  int  tipinde  deyerler saxlanilir   ve  hemin yerde  int tipinde  bir nece deyer saxlanila biler
	buna gorede bele dusne bilerik  bizim bir kucemiz var  orda int tipi olan adamlar qalir ama bircoxda ev var  ve ferqli ferqli unvadadir 4
	o evler yeni  int  kucesinde  bir nece  bir cox int saxlanila biler  adresleri ferqli ola biler  bu yadasin nece islemesine dair
	bir misaldi sadece
	demeli  biz burdan oyrenceymiz sey  deyerler  yadasda   bir  deyeri birde  adresi saxlanilir  her  ikisi ile  orada  ne olsudgunu
	gore  bilirik
*/
package main

import (
	"fmt"
)

func main() {
	// int tipinde deyisken yardip ona deyer teyin edirik
	var a int = 45
	fmt.Println("a nin deyeri:", a)
	//   indi bizim & ampersand adlandirdigimiz bu operatorla o deyerin adresini  almis olacagiq bunu ucun bu operatoru
	// o deyerin  qarsisina yazmagimiz  bes  edir
	fmt.Println("a nin  adresi:", &a) // cap edek a nin  adresi: 0xc0000ae000  bele bir cixis olacaq onaltiliq yazi sisteminde  cap olunur
	// biz  hemcinin bir deyerin adresini basqa bir deyiskene teyin ede  bilerik
	y := &a
	fmt.Println("y nin deyeri", y) //y nin deyeri 0xc000012140  gordymuz kimi her iksi eyni adresi tutur   burdanada  anlasildigi uzere
	// biz  y uzerinden de  x in deyerine  elaqe yarda  ve   deyerini deysdire bilerik
	// biz * opertaoru ile  her hansisa  adresde olan deyere baxa biler ve  onu yazdira bileriy
	fmt.Println("y deyeri", *y) //  unvan saxlayirdi ozunde hansiki yuxarda gosterdiyimiz kimi  a  unvanini  ama indi o unvandaki deyeri
	// bu operator vasitesi ile yazdira bileriy
	// biz y x -in  adresini saxladigi ucun y deyer teyin etsek  ikside eyni unvani gosterdiyi  ucun burada x da deyisecek bu
	// bir nov referansa benzeyir her iksi eyni yere referans edir ferz ede bilerik  yeni menim ve dosdumun elinde  bir adamin
	// unvani var ama eyni adresdi  biz o admi tapmaq isdesek ikmzde  eyni eve getmis olarqi ele deyilmi   indi bunu edek
	*y = 100 //  biz  y deyer elave  etdiyimiz ucun  onu *y  bu formatda yaziriq  ki biz her hansisa bir deyeri alip bu y de yerlesen adrese
	// gonderey
	fmt.Println("y in deyeri ", *y) //100   her ikisinin deyerini 100 oldugunu goruruk y deyisdirmeymiz bes etdiki  a  deyissin
	fmt.Println("y in deyeri ", a)  // 100  BUNU DAHA  COX  detalli sekilde  funksiallar  formatinda yazcam
	// *int    bu yeni  int   tipindeki deyerlerin adresini saxlayan  deyisken de  deye  bilerik

}
