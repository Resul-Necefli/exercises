package main

import "fmt"

type sheping interface {
	ecar() float64
}

/*
polimorfizim cox bicmlilikde  deyilir yeni eyni metod muxtelif novler uzerinde insa edildiyi ucun o muxtelif hemen novlerden
deyer qebul ederek herdefe  muxtelif novlerin cavbini bize  qaytarir   yeni  bele dusnek  bizim hava limanina gedirik
ve orada   bizden bilet isteyirler  ve bizde biletmizi gosteririk bizi buraxillar  o biletin basqalarindada olmasi onemli deyil
eyni biletle bir  cox  insan kecir  yalniz  o  biletin uzerinde  indi  kimin adi soyadi  varsa  o  adam kecir novbeti adamin biletide menimkinendi
ama  ferqli ad  soyad   yeni polimorfizimde  eyni metoddan bir cox  deyerleri cagirma olarin uzerinde  emeliyatlar  etmeye deyeilir
biz meselen    10  ferqli structumuz  var  ve  bu structlarin hamsinin  area()  adli  metodu var   biz  areani  1 nomreli yeni 1.area()
kimi  cagrsaq 1 nomrelinin adin soyadin  yeni  o sturutun  deyerlerin vercek  ve ya 2.area() ile  cagirsaq  2 cinin
*/
func interfacefunc(angless ...sheping) {

	for _, v := range angless {

		fmt.Println(v.ecar())
	}

}

type angle3 struct {
	h float64
	a float64
}

type angle4 struct {
	b float64
	c float64
}

// metodumuzu yarddaq
func (a angle3) ecar() float64 {
	return a.a * a.h
}

// eyni adli metoddan yene istfade  edrik
func (a angle4) ecar() float64 {
	return a.b * a.c
}

func main() {
	add := angle3{
		h: 3.5,
		a: 4.2,
	}
	aff := angle4{
		b: 5,
		c: 6,
	}
	interfacefunc(add, aff)

}

/*
polmorfizimin yaxsi terefi odurki meslen menim iki structum var ve  onlar  eyni iterface referans edirler amma  hersinin bir metodu var  ve  ferqli
menim  bu interface islemeycek cunki o metodun   herbirinin   bu structlarin    herbirnde  2 sinde olmasini isdycek yen
orxan ve resul  struct var  bularinda add()   sdd()  metodlari   burada  bu metodun herkisi hem  resulda   hemde orxanda  olmalidir polimorfizim bunun onune kecir
*/
